package fnfile

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

// GoSpec is a step that simplifies the use of running a "cli" that you make and code yourself
// you can store just the .go files, and compile your cli on-demand (similar to Magefile)
// but with the additional benefit of compiling _only_ when the generated executable is out-of-date
// as compared to the source files
type GoSpec struct {
	StepMeta
	Do Do `json:"do,omitempty"`
}

func (spec *GoSpec) UnmarshalJSON(data []byte) (err error) {
	*spec, err = UnmarshalGoSpec(data)
	return
}

func (spec GoSpec) Handle(w ResponseWriter, _ *FnContext) {

}

func UnmarshalGoSpecStep(data []byte) (Step, error) {
	return UnmarshalDefer(data)
}

func UnmarshalGoSpec(data []byte) (GoSpec, error) {
	var attemptErrs *multierror.Error

	// most steps (including this one) can be shortcut represented by a string
	// defer will unmarshal to a nested Sh
	sh, err := UnmarshalSh(data)
	if err == nil {
		return GoSpec{
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
		return GoSpec{
			Do: Do{
				Steps: steps,
			},
		}, nil
	}
	attemptErrs = multierror.Append(attemptErrs, fmt.Errorf("trying array shortcut: %w", err))

	type GoSpecAlias GoSpec
	var tmpGoSpec GoSpecAlias

	err = json.Unmarshal(data, &tmpGoSpec)
	if err == nil {
		return GoSpec(tmpGoSpec), nil
	}
	attemptErrs = multierror.Append(attemptErrs, fmt.Errorf("trying as Defer proper: %w", err))

	return GoSpec{}, GivingUp(attemptErrs)
}
