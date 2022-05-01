package fnfile

import (
	"context"
)

type Sh struct {
	StepCommon

	// Run defines a shell command.
	Run string `json:"run,omitempty"`

	// Dir is the desired working directory in which this command should execute in.
	Dir string `json:"dir,omitempty"`
}

func (sh *Sh) Visit(ctx context.Context, v StepVisitor) error {
	return v.VisitSh(ctx, sh)
}
