package fnfile

import (
	"encoding/json"
	"fmt"
)

type FnStepSpec struct {
	StepMeta
	Fn string `json:"fn,omitempty"`
}

func (fn *FnStepSpec) UnmarshalJSON(data []byte) (err error) {
	*fn, err = UnmarshalFnStep(data)
	return
}

func (fn FnStepSpec) Handle(w ResponseWriter, c *FnContext) {
	fnfile := c.FnFile()
	if target, ok := fnfile.Fns[fn.Fn]; ok {
		target.Do.Handle(w, c)
	}
}

func UnmarshalFnStepStep(data []byte) (Step, error) {
	return UnmarshalFnStep(data)
}

func UnmarshalFnStepFromStringData(data []byte) (FnStepSpec, error) {
	// most steps (including this one) can be shortcut represented by a string
	var tmpString string
	err := json.Unmarshal(data, &tmpString)
	if err != nil {
		return FnStepSpec{}, err
	}

	return FnStepSpec{
		Fn: tmpString,
	}, nil
}

func UnmarshalFnStep(data []byte) (FnStepSpec, error) {
	if fnStep, err := UnmarshalFnStepFromStringData(data); err == nil {
		return fnStep, nil
	}

	type FnStepSpecAlias FnStepSpec
	var tmpFnStepSpec FnStepSpecAlias

	err := json.Unmarshal(data, &tmpFnStepSpec)
	if err != nil {
		return FnStepSpec{}, fmt.Errorf("unmarshalling to FnStep proper: %w", err)
	}

	return FnStepSpec(tmpFnStepSpec), nil
}
