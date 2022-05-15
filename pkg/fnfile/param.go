package fnfile

import (
	"encoding/json"
)

type Params map[string]Param

func (p *Params) UnmarshalJSON(data []byte) error {
	type ParamsAlias Params
	var tmpParams ParamsAlias

	err := json.Unmarshal(data, &tmpParams)
	if err != nil {
		return err
	}

	// TODO fix double loop
	// we are looping over this twice
	// once here, once in validate, feels bad, maybe it doesn't matter
	for k, v := range tmpParams {
		nv := v
		nv.Name = k
		tmpParams[k] = nv
	}

	*p = Params(tmpParams)
	return nil
}

type Param Variable

func ParamFromString(val string) Param {
	return Param(VariableFromString(val))
}

func (p *Param) UnmarshalJSON(data []byte) error {
	newV, err := UnmarshalJSONTryAsString(data, ParamFromString)
	if err != nil {
		return err
	}

	*p = newV
	return nil
}
