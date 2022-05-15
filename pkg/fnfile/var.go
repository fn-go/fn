package fnfile

import (
	"encoding/json"
)

type Vars map[string]Variable

func (v *Vars) UnmarshalJSON(data []byte) error {
	type TmpVars Vars
	var tmpVars TmpVars

	err := json.Unmarshal(data, &tmpVars)
	if err != nil {
		return err
	}

	// TODO fix double loop
	// we are looping over this twice
	// once here, once in validate, feels bad, maybe it doesn't matter
	for k, v := range tmpVars {
		nv := v
		nv.name = k
		tmpVars[k] = nv
	}

	*v = Vars(tmpVars)
	return nil
}

type Variable struct {
	Name   string `json:"-"`
	Static string `json:"static,omitempty"`
	Tpl    string `json:"tpl,omitempty"`
	Sh     Sh     `json:"sh,omitempty"`
}

func VariableFromString(val string) Variable {
	return Variable{
		Tpl: val,
	}
}

func (v *Variable) UnmarshalJSON(data []byte) error {
	newV, err := UnmarshalJSONTryAsString(data, VariableFromString)
	if err != nil {
		return err
	}

	*v = newV
	return nil
}
