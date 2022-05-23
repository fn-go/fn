package fnfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/ghostsquad/go-timejumper"
)

type StepType string

type Steps []Step

func (steps *Steps) UnmarshalJSON(data []byte) error {
	// a series of steps can be shortcut represented as a single step
	step, err := StepFromJson(data)
	if err == nil {
		*steps = Steps{
			step,
		}
		return nil
	}

	// Otherwise unmarshal into a list of steps
	var sRaw []json.RawMessage
	err = json.Unmarshal(data, &sRaw)
	if err != nil {
		return fmt.Errorf("unmarshalling to []step: %w", err)
	}

	tmpSteps := make(Steps, len(sRaw))
	i := 0
	for _, s := range sRaw {
		step, err := StepFromJson(s)
		if err != nil {
			return err
		}
		tmpSteps[i] = step
		i++
	}

	*steps = tmpSteps
	return nil
}

type StepJson struct {
	Sh       *Sh        `json:"sh,omitempty"`
	Do       *Do        `json:"do,omitempty"`
	Parallel *Parallel  `json:"parallel,omitempty"`
	Defer    *DeferSpec `json:"defer,omitempty"`
}

// StepFromJson allows us to request a map-like syntax when unmarshalling
// But will resolve to the first field that's not nil
func StepFromJson(data []byte) (Step, error) {
	stepJson := StepJson{}
	err := json.Unmarshal(data, &stepJson)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(stepJson)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		val := field.Interface().(Step)
		if val != nil {
			return val, nil
		}
	}

	return nil, errors.New("yaml schema doesn't match. check the documentation for help")
}

type ProgrammaticNamer interface {
	ProgrammaticName() string
}

type Step interface {
	ProgrammaticNamer

	Exec(ResponseWriter, *CallInfo)
}

type StepMeta struct {
	// Name is the user defined name of the step. Defaults to "anonymous"
	Name string `json:"-"`
	// Locals are variables available to this step and all child steps (like a closure)
	Locals Vars `json:"vars,omitempty"`
	// Timeout is the bounding time limit (duration) for  before signalling for termination
	Timeout Duration `json:"timeout,omitempty"`

	clock    timejumper.Clock
	stepType StepType
	parent   ProgrammaticNamer
}

func (s StepMeta) ProgrammaticName() string {
	var prefix string

	if s.parent != nil {
		prefix = s.parent.ProgrammaticName() + ": "
	}

	return prefix + string(s.stepType) + ": " + s.Name
}

func NewStepMeta(stepType StepType, parent ProgrammaticNamer) StepMeta {
	m := StepMeta{
		Name:     "anonymous",
		Locals:   make(Vars),
		clock:    timejumper.RealClock{},
		stepType: stepType,
		parent:   parent,
	}

	return m
}
