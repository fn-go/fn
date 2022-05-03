package engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ghostsquad/go-timejumper"
	"github.com/hashicorp/go-multierror"

	"github.com/go-fn/fn/pkg/fnfile"
)

type DefaultStepHandler struct {
	cancel    context.CancelFunc
	mu        sync.Mutex
	Deferrals []*fnfile.DeferSpec
	time      timejumper.Clock
}

func (h *DefaultStepHandler) Handle(parentCtx context.Context, s fnfile.Step) error {
	var ctx context.Context
	ctx, h.cancel = context.WithCancel(parentCtx)

	var mErr *multierror.Error
	defer func() {
		for _, d := range h.Deferrals {
			select {
			case <-ctx.Done():
				mErr = multierror.Append(mErr, NewCancelledStepError(s.GetName(), ctx.Err()))
			default:
				mErr = multierror.Append(mErr, d.Visit(ctx, h))
			}
		}
	}()

	mErr = multierror.Append(mErr, s.Visit(ctx, h))
	return mErr.ErrorOrNil()
}

func (h *DefaultStepHandler) VisitDo(ctx context.Context, do *fnfile.Do) error {
	var mErr *multierror.Error

	for _, s := range do.Steps {
		select {
		case <-ctx.Done():
			mErr = multierror.Append(mErr, NewCancelledStepError(s.GetName(), ctx.Err()))
		default:
			mErr = multierror.Append(mErr, s.Visit(ctx, h))
		}
	}

	return mErr.ErrorOrNil()
}

func (h *DefaultStepHandler) VisitParallel(ctx context.Context, parallel *fnfile.Parallel) error {
	mErr := NewThreadSafeMultiError()

	for _, s := range parallel.Steps {
		go func(s fnfile.Step) {
			mErr.Append(s.Visit(ctx, h))
		}(s)
	}

	return mErr.ErrorOrNil()
}

func (h *DefaultStepHandler) VisitTry(ctx context.Context, try *fnfile.Try) error {
	err := try.Visit(ctx, h)
	if err != nil {
		fmt.Println("DEBUG: tried and failed!")
	}
	return nil
}

func (h *DefaultStepHandler) VisitSh(parentContext context.Context, sh *fnfile.Sh) error {
	ctx := context.Background()

	if sh.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(parentContext, time.Now().Add(time.Duration(sh.Timeout)))
		// Even though ctx will be expired, it is good practice to call its
		// cancellation function in any case. Failure to do so may keep the
		// context and its parent alive longer than necessary.
		defer cancel()
	}

	select {
	case <-ctx.Done():
	default:
		fmt.Println(sh.Run)
	}

	return nil
}

func (h *DefaultStepHandler) VisitDefer(_ context.Context, spec *fnfile.DeferSpec) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.Deferrals = append(h.Deferrals, spec)
	return nil
}

func (h *DefaultStepHandler) VisitReturn(_ context.Context, _ *fnfile.ReturnSpec) error {
	h.cancel()
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
