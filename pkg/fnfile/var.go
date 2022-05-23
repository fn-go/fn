package fnfile

type Vars map[string]Variable

func (v *Vars) UnmarshalJSON(data []byte) error {
	tmpVal := make(Vars)

	err := UnmarshalJSONToNamedMap(tmpVal, func(name string) Variable {
		return Variable{
			Name: name,
		}
	}, data)
	if err != nil {
		return err
	}

	*v = tmpVal
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
