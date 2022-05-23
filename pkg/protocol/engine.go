package protocol

import (
	"bytes"
	"context"
	"sync"

	"github.com/go-fn/fn/pkg/fnfile"
)

type Engine struct {
	cancel context.CancelFunc
	mu     sync.Mutex
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
		writer: opts.Writer,
	}

	return &eng, nil
}

func (g *Engine) Run(parentCtx context.Context, fn fnfile.Fn) error {
	var ctx context.Context
	ctx, g.cancel = context.WithCancel(parentCtx)

	in := &bytes.Buffer{}

	fn.Do.Exec(g.writer, fnfile.NewCallInfo(ctx, in))
	return g.writer.ErrorOrNil()
}
