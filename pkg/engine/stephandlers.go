package engine

import (
	"context"

	"github.com/go-fn/fn/pkg/fnfile"
)

type DefaultStepHandler struct {
	Deferrals []*fnfile.DeferSpec
}

func (h *DefaultStepHandler) Handle(ctx context.Context, s fnfile.Step) error {
	defer func() {
		for _, d := range h.Deferrals {
			// TODO handle errors here
			_ = d.Visit(ctx, h)
		}
	}()
	return s.Visit(ctx, h)
}

func (h *DefaultStepHandler) VisitDo(ctx context.Context, do *fnfile.Do) error {
	for _, s := range do.Steps {
		return s.Visit(ctx, h)
	}

	return nil
}

func (h *DefaultStepHandler) VisitParallel(ctx context.Context, parallel *fnfile.Parallel) error {
	for _, s := range parallel.Steps {
		go func(s fnfile.Step) {
			_ := s.Visit(ctx, h)
		}(s)
	}

	return nil
}

func (h *DefaultStepHandler) VisitTry(ctx context.Context, try *fnfile.Try) error {
	// Run try but ignore errors
	return nil
}

func (h *DefaultStepHandler) VisitSh(ctx context.Context, sh *fnfile.Sh) error {
	// Run sh & return sh errors
	return nil
}

func (h *DefaultStepHandler) VisitDefer(ctx context.Context, spec *fnfile.DeferSpec) error {
	h.Deferrals = append(h.Deferrals, spec)
	return nil
}

func (h *DefaultStepHandler) VisitReturn(cancel context.CancelFunc, spec *fnfile.ReturnSpec) error {
	cancel()
	return nil
}

func (h *DefaultStepHandler) VisitMatrix(ctx context.Context, matrix *fnfile.Matrix) error {
	// Do matrix things
	return nil
}

func (h *DefaultStepHandler) VisitWait(ctx context.Context, wait *fnfile.Wait) error {
	// Do Wait/Watch things
	return nil
}
