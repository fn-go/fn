package fnfile

import (
	"context"
	"encoding/json"
	"fmt"
)

type Do struct {
	StepMeta
	Steps Steps `json:"steps"`
}

func (do Do) Accept(visitor StepVisitor) {
	visitor.VisitDo(do)
}

func (do *Do) UnmarshalJSON(data []byte) (err error) {
	*do, err = UnmarshalToDo(data)
	return
}

func (do Do) Handle(w ResponseWriter, c *FnContext) {
	parentCtx := c.Context()
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	visitor := NewStepVisitor(w, c.CloneWith(ctx))

	for _, s := range do.Steps {
		if w.ErrorOrNil() != nil {
			cancel()
		}
		select {
		case <-parentCtx.Done():
			w.Error(ctx.Err())
			return
		default:
			s.Accept(visitor)
		}
	}
}

func UnmarshalToDo(data []byte) (Do, error) {
	// most steps (including this one) can be shortcut represented as a Sh step
	sh, err := UnmarshalSh(data)
	if err == nil {
		return Do{
			Steps: Steps{
				sh,
			},
		}, nil
	}

	type DoAlias Do
	var tmpDo DoAlias

	err = json.Unmarshal(data, &tmpDo)
	if err != nil {
		return Do{}, fmt.Errorf("unmarshalling to DoAlias: %w", err)
	}

	return Do(tmpDo), nil
}
