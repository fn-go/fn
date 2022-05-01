package fnfile

import (
	"context"
)

type Wait struct {
	StepCommon
}

func (w *Wait) Visit(ctx context.Context, v StepVisitor) error {
	return v.VisitWait(ctx, w)
}
