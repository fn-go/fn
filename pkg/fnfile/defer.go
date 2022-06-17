package fnfile

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

type DeferSpec struct {
	StepMeta
	Do Do `json:"do,omitempty"`
}

func (spec *DeferSpec) UnmarshalJSON(data []byte) (err error) {
	*spec, err = UnmarshalDefer(data)
	return
}

func (spec DeferSpec) Handle(w ResponseWriter, _ *FnContext) {
	w.Defer(StepHandlerFunc(spec.Do.Handle))
}

func UnmarshalDeferStep(data []byte) (Step, error) {
	return UnmarshalDefer(data)
}

func UnmarshalDefer(data []byte) (DeferSpec, error) {
	var attemptErrs *multierror.Error

	// most steps (including this one) can be shortcut represented by a string
	// defer will unmarshal to a nested Sh
	sh, err := UnmarshalSh(data)
	if err == nil {
		return DeferSpec{
			Do: Do{
				Steps: Steps{
					sh,
				},
			},
		}, nil
	}
	attemptErrs = multierror.Append(attemptErrs, fmt.Errorf("trying sh shortcut: %w", err))

	// this can also be shortcut represented as just an array
	steps, err := UnmarshalSteps(data)
	if err == nil {
		return DeferSpec{
			Do: Do{
				Steps: steps,
			},
		}, nil
	}
	attemptErrs = multierror.Append(attemptErrs, fmt.Errorf("trying array shortcut: %w", err))

	type DeferAlias DeferSpec
	var tmpDefer DeferAlias

	err = json.Unmarshal(data, &tmpDefer)
	if err == nil {
		return DeferSpec(tmpDefer), nil
	}
	attemptErrs = multierror.Append(attemptErrs, fmt.Errorf("trying as Defer proper: %w", err))

	return DeferSpec{}, GivingUp(attemptErrs)
}
