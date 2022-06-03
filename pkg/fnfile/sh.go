package fnfile

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/oklog/run"
	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Sh struct {
	StepMeta
	// Run defines a shell command.
	Run string `json:"run,omitempty"`
	// Dir is the desired working directory in which this command should execute in.
	Dir string `json:"dir,omitempty"`
}

func (sh Sh) Validate() error {
	var mErr *multierror.Error

	if sh.Run == "" {
		mErr = multierror.Append(mErr, fmt.Errorf("invalid run statement: %s", sh.Run))
	}

	return mErr.ErrorOrNil()
}

func (sh Sh) Accept(visitor StepVisitor) {
	visitor.VisitSh(sh)
}

func (sh Sh) Handle(w ResponseWriter, c *FnContext) {
	if err := sh.Validate(); err != nil {
		w.Error(fmt.Errorf("validating sh: %w", err))
		return
	}

	var deadline time.Time
	if sh.Timeout > 0 {
		deadline = c.clock.Now().Add(time.Duration(sh.Timeout))
	} else {
		deadline = c.clock.Now().Add(time.Minute * 2)
	}

	ctx, cancel := context.WithDeadline(c.Context(), deadline)
	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	r, err := interp.New(
		interp.Params("-e"),
		interp.Env(expand.ListEnviron(os.Environ()...)),
		interp.OpenHandler(openHandler),
		interp.StdIO(c.In(), w.OutWriter(), w.ErrOutWriter()),
		interp.Dir(sh.Dir),
	)

	if err != nil {
		w.Error(err)
		return
	}

	p, err := syntax.NewParser().Parse(strings.NewReader(sh.Run), "")
	if err != nil {
		w.Error(err)
		return
	}

	var group run.Group
	group.Add(func() error {
		err := r.Run(ctx, p)
		if err != nil {
			return fmt.Errorf("running sh [%s]: %w", sh.Run, err)
		}
		return nil
	}, func(err error) {
		if cancel != nil {
			cancel()
		}
	})

	group.Add(func() error {
		select {
		case <-c.Context().Done():
			// If we receive a SIGINT, let signal propagation handle the cancellation
			// This is to avoid the problem presented in:
			// TODO link to signal issue in mvdan/sh
			if !errors.Is(c.Context().Err(), run.SignalError{}) {
				cancel()
			}
		case <-ctx.Done():
			cancel()
		}
		return nil
	}, func(err error) {
		cancel()
	})

	err = group.Run()
	if err != nil {
		w.Error(err)
		return
	}
}

func (sh *Sh) UnmarshalJSON(data []byte) (err error) {
	*sh, err = UnmarshalSh(data)
	return
}

// UnmarshalSh attempts to unmarshal to a string to generate
// a simple Sh step (only the run field)
// otherwise attempts a full unmarshal to Sh step
// returns an error if neither is possible
func UnmarshalSh(data []byte) (Sh, error) {
	// a Sh step can be shortcut represented as just a string
	var tmpString string
	err := json.Unmarshal(data, &tmpString)
	if err == nil {
		tmpSh := Sh{
			Run: tmpString,
		}
		return tmpSh, nil
	}

	type ShAlias Sh
	var tmpSh ShAlias

	err = json.Unmarshal(data, &tmpSh)
	if err != nil {
		return Sh{}, err
	}

	return Sh(tmpSh), nil
}

func openHandler(ctx context.Context, path string, flag int, perm os.FileMode) (io.ReadWriteCloser, error) {
	if path == "/dev/null" {
		return devNull{}, nil
	}
	return interp.DefaultOpenHandler()(ctx, path, flag, perm)
}

type devNull struct{}

func (devNull) Read(_ []byte) (int, error)  { return 0, io.EOF }
func (devNull) Write(p []byte) (int, error) { return len(p), nil }
func (devNull) Close() error                { return nil }
