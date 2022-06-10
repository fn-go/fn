package fnfile

import (
	"encoding/json"
	"fmt"
)

type Dynamic struct {
	StepMeta
	JSON string `json:"json,omitempty"`

	//resolved Step
}

func (dynamic Dynamic) Handle(w ResponseWriter, c *FnContext) {

}

func (dynamic *Dynamic) UnmarshalJSON(data []byte) (err error) {
	*dynamic, err = UnmarshalDynamic(data)
	return
}

func UnmarshalDynamicStep(data []byte) (Step, error) {
	return UnmarshalDynamic(data)
}

func UnmarshalDynamic(data []byte) (Dynamic, error) {
	// most steps (including this one) can be shortcut represented by a string
	// defer will unmarshal to a nested Sh
	var tmpString string
	err := json.Unmarshal(data, &tmpString)
	if err == nil {
		return Dynamic{
			JSON: tmpString,
		}, nil
	}

	type DynamicAlias Dynamic
	var tmpDynamic DynamicAlias

	err = json.Unmarshal(data, &tmpDynamic)
	if err != nil {
		return Dynamic{}, fmt.Errorf("unmarshalling to Dynamic proper: %w", err)
	}

	return Dynamic(tmpDynamic), nil
}
