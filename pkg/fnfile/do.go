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

func (do *Do) UnmarshalJSON(data []byte) (err error) {
	*do, err = UnmarshalDo(data)
	return
}

func (do Do) Handle(w ResponseWriter, c *FnContext) {
	parentCtx := c.Context()
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	for _, s := range do.Steps {
		if w.ErrorOrNil() != nil {
			cancel()
		}
		select {
		case <-parentCtx.Done():
			w.Error(ctx.Err())
			return
		default:
			s.Handle(w, c.CloneWith(ctx))
		}
	}
}

func UnmarshalDoStep(data []byte) (Step, error) {
	return UnmarshalDo(data)
}

func UnmarshalDo(data []byte) (Do, error) {
	// most steps (including this one) can be shortcut represented by a string
	// Do will unmarshal to a nested Sh
	sh, err := UnmarshalSh(data)
	if err == nil {
		return Do{
			Steps: Steps{
				sh,
			},
		}, nil
	}

	// this can also be shortcut represented as just an array
	steps, err := UnmarshalSteps(data)
	if err == nil {
		return Do{
			Steps: steps,
		}, nil
	}

	type DoAlias Do
	var tmpDo DoAlias

	err = json.Unmarshal(data, &tmpDo)
	if err != nil {
		return Do{}, fmt.Errorf("unmarshalling to Do proper: %w", err)
	}

	return Do(tmpDo), nil
}
