package fnfile

import (
	"context"
)

type Parallel struct {
	StepCommon

	Steps    []Step `json:"steps"`
	FailFast bool   `json:"failFast"`
	Limit    int    `json:"limit"`
}

func (p *Parallel) Visit(ctx context.Context, v StepVisitor) error {
	return v.VisitParallel(ctx, p)
}
