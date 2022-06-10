package engine

import (
	"context"
	"sync"

	"github.com/go-fn/fn/pkg/fnfile"
)

type Engine struct {
	cancel context.CancelFunc
	mu     sync.Mutex
	writer fnfile.ResponseWriter
}

type Options struct {
	Writer fnfile.ResponseWriter
}

// New returns a new engine
func New(options ...func(engineOptions *Options)) *Engine {
	opts := &Options{
		Writer: fnfile.DiscardResponseWriter{},
	}

	for _, o := range options {
		o(opts)
	}

	eng := Engine{
		mu:     sync.Mutex{},
		writer: opts.Writer,
	}

	return &eng
}

func (g *Engine) Run(parentCtx context.Context, fn fnfile.FnDef) error {
	var ctx context.Context
	ctx, g.cancel = context.WithCancel(parentCtx)

	// TODO move all this into the `fn`
	callInfo := fnfile.NewCallInfo(ctx)
	w2, deferrals := fnfile.WithNewDeferrals(g.writer)

	defer func(deferrals *[]fnfile.StepHandler) {
		for i := len(*deferrals) - 1; i >= 0; i-- {
			(*deferrals)[i].Handle(w2, callInfo)
		}
	}(deferrals)

	fn.Do.Handle(w2, callInfo)

	return g.writer.ErrorOrNil()
}
