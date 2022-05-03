package fnfile

import (
	"context"
)

type Do struct {
	StepCommon

	Steps []Step `json:"steps"`
}

func (d *Do) Visit(ctx context.Context, v StepVisitor) error {
	return v.VisitDo(ctx, d)
}
