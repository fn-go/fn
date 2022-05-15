package protocol

import (
	"bytes"
	"context"
	"sync"

	"github.com/ghostsquad/go-timejumper"

	"github.com/go-fn/fn/pkg/fnfile"
)

type Engine struct {
	cancel context.CancelFunc
	mu     sync.Mutex
	m      map[string]Handler
	time   timejumper.Clock
	writer fnfile.ResponseWriter
}

type EngineOptions struct {
	Writer fnfile.ResponseWriter
}

// New returns a new engine
func New(options ...func(engineOptions *EngineOptions)) (*Engine, error) {
	opts := &EngineOptions{
		Writer: fnfile.DiscardResponseWriter{},
	}

	for _, o := range options {
		o(opts)
	}

	eng := Engine{
		mu:     sync.Mutex{},
		m:      make(map[string]Handler),
		time:   timejumper.RealClock{},
		writer: opts.Writer,
	}

	return &eng, nil
}

type Handler interface {
	Do(context.Context, *fnfile.Step)
}

type HandleFunc func(context.Context, *fnfile.Step) error

func (g *Engine) Handle(name string, handler Handler) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if name == "" {
		panic("invalid name")
	}
	if handler == nil {
		panic("invalid handler")
	}

	if g.m == nil {
		g.m = make(map[string]Handler)
	}

	g.m[name] = handler
}

func (g *Engine) Run(parentCtx context.Context, fn fnfile.Fn) error {
	var ctx context.Context
	ctx, g.cancel = context.WithCancel(parentCtx)

	in := &bytes.Buffer{}

	for _, s := range fn.Do {
		s.Exec(g.writer, fnfile.NewCallInfo(ctx, in))
		err := g.writer.ErrorOrNil()
		if err != nil {
			return err
		}
	}

	return nil
}
