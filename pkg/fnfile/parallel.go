package fnfile

import (
	"context"

	"github.com/oklog/run"
)

type Parallel struct {
	StepMeta

	Steps    []Step `json:"steps"`
	FailFast bool   `json:"failFast"`
	Limit    int    `json:"limit"`
}

func (p Parallel) Exec(w ResponseWriter, c *CallInfo) {
	validateHandlerParams(w, c)
	if p.FailFast {
		p.failFast(w, c)
		return
	}

	for _, s := range p.Steps {
		go s.Exec(w, c)
	}
}

func (p Parallel) failFast(w ResponseWriter, c *CallInfo) {
	ctx, cancel := context.WithCancel(c.Context())
	defer cancel()

	var group run.Group
	for _, s := range p.Steps {
		group.Add(func() error {
			w2, _ := withNewDeferrals(w)
			s.Exec(w2, c.CloneWith(ctx))
			return w.ErrorOrNil()
		}, func(err error) {
			cancel()
		})
	}

	w.Error(group.Run())
}
