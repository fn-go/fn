package fnfile

import (
	"context"
)

type Steps []string

// StepVisitor is part of the Decoupled Visitor Pattern
// https://making.pusher.com/alternatives-to-sum-types-in-go/
type StepVisitor interface {
	VisitDo(ctx context.Context, do *Do) error
	VisitParallel(ctx context.Context, parallel *Parallel) error
	VisitTry(ctx context.Context, try *Try) error
	VisitSh(ctx context.Context, sh *Sh) error
	VisitDefer(ctx context.Context, spec *DeferSpec) error
	VisitReturn(ctx context.Context, spec *ReturnSpec) error
	VisitMatrix(ctx context.Context, matrix *Matrix) error
	VisitWait(ctx context.Context, wait *Wait) error
}

type Step interface {
	GetName() string
	Visit(ctx context.Context, v StepVisitor) error
}

// StepCommon is a specific command (or even another task) to execute
// The use of Run & Args is mutually exclusive with Task
type StepCommon struct {
	Step

	Name string `json:"name"`

	Locals *Vars `json:"vars,omitempty"`

	// Timeout is the umbrella bounding time limit (duration) for the task before signalling for termination via SIGINT.
	Timeout Duration `json:"timeout,omitempty"`

	// GracefulTermination is the bounding time limit (duration) for this task before sending subprocesses a SIGKILL.
	GracefulTermination Duration `json:"gracefulTermination,omitempty"`
}
