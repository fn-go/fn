package fnfile

import (
	"context"
)

// ReturnSpec is a step hook which allows a fn to return early (skipping subsequent steps) just like
// a `return` statement in Go.
type ReturnSpec struct {
	StepCommon
}

func (r *ReturnSpec) Visit(ctx context.Context, v StepVisitor) error {
	return v.VisitReturn(ctx, r)
}
