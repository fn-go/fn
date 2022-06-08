package fnfile

import (
	"encoding/json"
	"fmt"
)

type DeferSpec struct {
	StepMeta
	Do
}

func (spec *DeferSpec) UnmarshalJSON(data []byte) (err error) {
	*spec, err = UnmarshalDefer(data)
	return
}

func (spec DeferSpec) Handle(w ResponseWriter, _ *FnContext) {
	w.Defer(StepHandlerFunc(spec.Do.Handle))
}

func UnmarshalDeferStep(data []byte) (Step, error) {
	return UnmarshalDefer(data)
}

func UnmarshalDefer(data []byte) (DeferSpec, error) {
	// most steps (including this one) can be shortcut represented by a string
	// defer will unmarshal to a nested Sh
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

	// this can also be shortcut represented as just an array
	steps, err := UnmarshalSteps(data)
	if err == nil {
		return DeferSpec{
			Do: Do{
				Steps: steps,
			},
		}, nil
	}

	type DeferAlias DeferSpec
	var tmpDefer DeferAlias

	err = json.Unmarshal(data, &tmpDefer)
	if err != nil {
		return DeferSpec{}, fmt.Errorf("unmarshalling to Defer proper: %w", err)
	}

	return DeferSpec(tmpDefer), nil
}
