package fnfile

import (
	"context"
)

type Try struct {
	StepCommon
}

func (t *Try) Visit(ctx context.Context, v StepVisitor) error {
	return v.VisitTry(ctx, t)
}
