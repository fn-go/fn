package fnfile

import (
	"encoding/json"

	"github.com/ghostsquad/go-timejumper"
)

type Steps []Step

type Step interface {
	Exec(ResponseWriter, *CallInfo)
}

// StepCommon is a specific command (or even another task) to execute
// The use of Run & Args is mutually exclusive with Task
type StepCommon struct {
	Name   string `json:"name"`
	Locals Vars   `json:"vars,omitempty"`
	// Timeout is the bounding time limit (duration) for  before signalling for termination
	Timeout Duration `json:"timeout,omitempty"`
	clock   timejumper.Clock
}

type StepCommonOptions struct {
	Locals  Vars
	Timeout Duration
	Clock   timejumper.Clock
}

type StepCommonOption func(options *StepCommonOptions)

func NewStepCommon(name string, options ...StepCommonOption) StepCommon {
	opts := &StepCommonOptions{
		Locals: make(Vars),
		Clock:  timejumper.RealClock{},
	}
	for _, o := range options {
		o(opts)
	}

	return StepCommon{
		Name:    name,
		Locals:  opts.Locals,
		clock:   opts.Clock,
		Timeout: opts.Timeout,
	}
}

func (s *StepCommon) UnmarshalJSON(data []byte) error {
	type StepCommonAlias StepCommon
	var tmpStepCommon StepCommonAlias

	err := json.Unmarshal(data, &tmpStepCommon)
	if err != nil {
		return err
	}

	if tmpStepCommon.Locals == nil {
		tmpStepCommon.Locals = make(Vars)
	}

	tmpStepCommon.clock = timejumper.RealClock{}

	*s = StepCommon(tmpStepCommon)
	return nil
}
