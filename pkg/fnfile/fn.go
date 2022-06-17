package fnfile

import (
	"encoding/json"
	"fmt"
)

type Fn struct {
	StepMeta
	Call string `json:"call,omitempty"`
}

func (fn *Fn) UnmarshalJSON(data []byte) (err error) {
	*fn, err = UnmarshalFn(data)
	return
}

func (fn Fn) Handle(w ResponseWriter, c *FnContext) {
	fnfile := c.FnFile()
	if target, ok := fnfile.Fns[fn.Call]; ok {
		target.Do.Handle(w, c)
	}
}

func UnmarshalFnStep(data []byte) (Step, error) {
	return UnmarshalFn(data)
}

func UnmarshalFnFromStringData(data []byte) (Fn, error) {
	// most steps (including this one) can be shortcut represented by a string
	var tmpString string
	err := json.Unmarshal(data, &tmpString)
	if err != nil {
		return Fn{}, err
	}

	return Fn{
		Call: tmpString,
	}, nil
}

func UnmarshalFn(data []byte) (Fn, error) {
	if fnStep, err := UnmarshalFnFromStringData(data); err == nil {
		return fnStep, nil
	}

	type FnStepSpecAlias Fn
	var tmpFnStepSpec FnStepSpecAlias

	err := json.Unmarshal(data, &tmpFnStepSpec)
	if err != nil {
		return Fn{}, fmt.Errorf("unmarshalling to Name proper: %w", err)
	}

	return Fn(tmpFnStepSpec), nil
}
