package fnfile

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

type Fns []Fn

func (fns *Fns) UnmarshalJSON(data []byte) (err error) {
	*fns, err = UnmarshalFns(data)
	return
}

func UnmarshalFns(data []byte) (Fns, error) {
	var attemptErrs *multierror.Error

	// most steps (including this one) can be shortcut represented by a string
	fn, err := UnmarshalFnFromStringData(data)
	if err == nil {
		return Fns{
			fn,
		}, nil
	}
	attemptErrs = multierror.Append(attemptErrs, fmt.Errorf("trying string shortcut: %w", err))

	type FnsAlias Fns
	var tmpFns FnsAlias

	err = json.Unmarshal(data, &tmpFns)
	if err == nil {
		return Fns(tmpFns), nil
	}
	attemptErrs = multierror.Append(attemptErrs, fmt.Errorf("trying as Fns proper: %w", err))

	return Fns{}, GivingUp(attemptErrs)
}
