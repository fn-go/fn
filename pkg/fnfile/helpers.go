package fnfile

import (
	"encoding/json"
)

type fromString[T any] func(val string) T

func UnmarshalJSONTryAsString[T any](data []byte, fromFn fromString[T]) (T, error) {
	var strVal string
	err := json.Unmarshal(data, &strVal)
	if err == nil {
		return fromFn(strVal), nil
	}

	var newT T
	err = json.Unmarshal(data, &newT)
	return newT, err
}

func UnmarshalJSONToNamedMap[T any](m map[string]T, newFn func(name string) T, data []byte) error {
	mRaw := make(map[string]json.RawMessage)

	err := json.Unmarshal(data, &mRaw)
	if err != nil {
		return err
	}

	for k, v := range mRaw {
		tmpVal := newFn(k)
		err := json.Unmarshal(v, &tmpVal)
		if err != nil {
			return err
		}

		m[k] = tmpVal
	}

	return nil
}
