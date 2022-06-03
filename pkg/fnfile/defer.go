package fnfile

import (
	"encoding/json"
	"fmt"
)

type DeferSpec struct {
	StepMeta
	Do
}

func (spec DeferSpec) Accept(visitor StepVisitor) {
	visitor.VisitDefer(spec)
}

func (spec DeferSpec) Handle(w ResponseWriter, _ *FnContext) {
	w.Defer(StepHandlerFunc(spec.Do.Handle))
}

func UnmarshalDefer(data []byte) (DeferSpec, error) {
	// most steps (including this one) can be shortcut represented as a Sh step
	sh, err := UnmarshalSh(data)
	if err == nil {
		return DeferSpec{
			Do: Do{
				Steps: Steps{
					sh,
				},
			},
		}, nil
	}

	type DeferAlias DeferSpec
	var tmpDefer DeferAlias

	err = json.Unmarshal(data, &tmpDefer)
	if err != nil {
		return DeferSpec{}, fmt.Errorf("unmarshalling to DeferAlias: %w", err)
	}

	return DeferSpec(tmpDefer), nil
}
