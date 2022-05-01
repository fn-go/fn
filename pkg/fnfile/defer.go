package fnfile

import (
	"context"
)

type DeferSpec struct {
	StepCommon
}

func (d *DeferSpec) Visit(ctx context.Context, v StepVisitor) error {
	return v.VisitDefer(ctx, d)
}
