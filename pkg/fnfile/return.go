package fnfile

import (
	"encoding/json"
)

// ReturnSpec is a step hook which allows a fn to return early (skipping subsequent steps) just like
// a `return` statement in Go.
type ReturnSpec struct {
	StepMeta
}

func (r *ReturnSpec) UnmarshalJSON(data []byte) (err error) {
	*r, err = UnmarshalReturn(data)
	return
}

func (r ReturnSpec) Handle(_ ResponseWriter, c *FnContext) {
	c.Context().Done()
}

func UnmarshalReturnStep(data []byte) (Step, error) {
	return UnmarshalReturn(data)
}

func UnmarshalReturn(data []byte) (ReturnSpec, error) {
	type ReturnSpecAlias ReturnSpec
	var tmpReturn ReturnSpecAlias

	err := json.Unmarshal(data, &tmpReturn)

	return ReturnSpec(tmpReturn), err
}
