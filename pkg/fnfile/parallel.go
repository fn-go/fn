package fnfile

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/oklog/run"
)

type Parallel struct {
	StepMeta

	Steps    Steps `json:"steps"`
	FailFast bool  `json:"failFast"`
	Limit    int   `json:"limit"`
}

func (p *Parallel) UnmarshalJSON(data []byte) (err error) {
	*p, err = UnmarshalParallel(data)
	return
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

	for _, s := range p.Steps {
		go s.Handle(w, localCallInfo)
	}
}

func (p Parallel) failFast(w ResponseWriter, c *FnContext) {
	parentCtx := c.Context()
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	localCallInfo := c.CloneWith(ctx)

	var group run.Group
	for _, s := range p.Steps {
		group.Add(func() error {
			s.Handle(w, localCallInfo)
			return w.ErrorOrNil()
		}, func(err error) {
			cancel()
		})
	}

	w.Error(group.Run())
}

func UnmarshalParallelStep(data []byte) (Step, error) {
	return UnmarshalParallel(data)
}

func UnmarshalParallel(data []byte) (Parallel, error) {
	// most steps (including this one) can be shortcut represented by a string
	// Parallel will unmarshal to a nested Sh
	sh, err := UnmarshalSh(data)
	if err == nil {
		return Parallel{
			Steps: Steps{
				sh,
			},
		}, nil
	}

	// this can also be shortcut represented as just an array
	steps, err := UnmarshalSteps(data)
	if err == nil {
		return Parallel{
			Steps: steps,
		}, nil
	}

	type ParallelAlias Parallel
	var tmpParallel ParallelAlias

	err = json.Unmarshal(data, &tmpParallel)
	if err != nil {
		return Parallel{}, fmt.Errorf("unmarshalling to Parallel proper: %w", err)
	}

	return Parallel(tmpParallel), nil
}
