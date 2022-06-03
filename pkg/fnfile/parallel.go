package fnfile

import (
	"context"

	"github.com/oklog/run"
)

type Parallel struct {
	StepMeta

	Steps    Steps `json:"steps"`
	FailFast bool  `json:"failFast"`
	Limit    int   `json:"limit"`
}

func (p Parallel) Accept(visitor StepVisitor) {
	visitor.VisitParallel(p)
}

func (p Parallel) Handle(w ResponseWriter, c *FnContext) {
	if p.FailFast {
		p.failFast(w, c)
	} else {
		p.bestEffort(w, c)
	}
}

func (p Parallel) bestEffort(w ResponseWriter, c *FnContext) {
	parentCtx := c.Context()
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	localCallInfo := c.CloneWith(ctx)

	visitor := NewStepVisitor(w, localCallInfo)

	for _, s := range p.Steps {
		go s.Accept(visitor)
	}
}

func (p Parallel) failFast(w ResponseWriter, c *FnContext) {
	parentCtx := c.Context()
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	localCallInfo := c.CloneWith(ctx)

	visitor := NewStepVisitor(w, localCallInfo)

	var group run.Group
	for _, s := range p.Steps {
		group.Add(func() error {
			s.Accept(visitor)
			return w.ErrorOrNil()
		}, func(err error) {
			cancel()
		})
	}

	w.Error(group.Run())
}

func UnmarshalParallel(data []byte) (Parallel, error) {
	return Parallel{}, nil
}
