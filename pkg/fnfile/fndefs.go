package fnfile

type FnDefs map[string]FnDef

func (f *FnDefs) UnmarshalJSON(data []byte) (err error) {
	*f, err = UnmarshalFnDefs(data)
	return
}

func UnmarshalFnDefs(data []byte) (FnDefs, error) {
	tmpVal := make(FnDefs)

	err := UnmarshalJSONToNamedMap(tmpVal, func(name string) FnDef {
		return FnDef{
			Name: name,
		}
	}, data)

	if err != nil {
		return FnDefs{}, err
	}

	return tmpVal, nil
}
