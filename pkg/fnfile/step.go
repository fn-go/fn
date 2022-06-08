package fnfile

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/samber/lo"
)

type Steps []Step

func (steps *Steps) UnmarshalJSON(data []byte) error {
	result, err := UnmarshalSteps(data)
	if err != nil {
		return err
	}
	*steps = result
	return nil
}

func UnmarshalSteps(data []byte) (Steps, error) {
	var tmpList []json.RawMessage
	err := json.Unmarshal(data, &tmpList)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling steps to list: %w", err)
	}

	var mErr *multierror.Error

	tmpSteps := make([]Step, len(tmpList))
	for i, s := range tmpList {
		tmpSteps[i], err = UnmarshalStep(s)
		mErr = multierror.Append(mErr, err)
	}

	err = mErr.ErrorOrNil()
	if err != nil {
		return nil, err
	}

	return tmpSteps, nil
}

type Step interface {
	Handle(w ResponseWriter, c *FnContext)
}

type StepVisitor struct {
	VisitDefer    func(d DeferSpec)
	VisitDo       func(do Do)
	VisitFnStep   func(fn FnStepSpec)
	VisitMatrix   func(m Matrix)
	VisitParallel func(p Parallel)
	VisitReturn   func(r ReturnSpec)
	VisitSh       func(sh Sh)
}

type StepMeta struct {
	// Name is the user defined name of the step. Defaults to "anonymous"
	Name string `json:"-"`
	// Locals are variables available to this step and all child steps (like a closure)
	Locals Vars `json:"vars,omitempty"`
	// Timeout is the bounding time limit (duration) for  before signaling for termination
	Timeout Duration `json:"timeout,omitempty"`
}

type StepName string

const (
	DeferStep    StepName = "defer"
	DoStep       StepName = "do"
	FnStep       StepName = "fn"
	MatrixStep   StepName = "matrix"
	ParallelStep StepName = "parallel"
	ReturnStep   StepName = "return"
	ShStep       StepName = "sh"
	DynamicStep  StepName = "dynamic"
)

type StepUnmarshaller func(data []byte) (Step, error)

type stepUnmarshalTuple struct {
	name StepName
	fn   StepUnmarshaller
}

var unmarshalPriorities []stepUnmarshalTuple

var unmarshalFuncsMap = map[StepName]StepUnmarshaller{}

func init() {
	unmarshalPriorities = []stepUnmarshalTuple{
		{
			name: ShStep,
			fn:   UnmarshalShStep,
		},
		{
			name: DoStep,
			fn:   UnmarshalDoStep,
		},
		{
			name: ParallelStep,
			fn:   UnmarshalParallelStep,
		},
		{
			name: FnStep,
			fn:   UnmarshalFnStepStep,
		},
		{
			name: MatrixStep,
			fn:   UnmarshalMatrixStep,
		},
		{
			name: DeferStep,
			fn:   UnmarshalDeferStep,
		},
		{
			name: ReturnStep,
			fn:   UnmarshalReturnStep,
		},
		{
			name: DynamicStep,
			fn:   UnmarshalDynamicStep,
		},
	}

	for _, v := range unmarshalPriorities {
		unmarshalFuncsMap[v.name] = v.fn
	}
}

func UnmarshalStep(data []byte) (Step, error) {
	// check if the object is "keyed" first aka:
	// "sh": {}
	step, err := UnmarshalKeyedStep(data)
	if err == nil {
		return step, nil
	}

	// if we are dealing with an unkeyed object, let's try to unmarshal to various other types,
	// starting with the most common/expected types first
	// in the case that multiple steps have identical fields, this becomes a priority list
	// of which step wins in the face of ambiguity
	for _, tuple := range unmarshalPriorities {
		step, err := tuple.fn(data)
		if err == nil {
			return step, nil
		}
	}

	return nil, errors.New("unknown data, could not unmarshal to any step type")
}

func UnmarshalKeyedStep(data []byte) (Step, error) {
	tmpMap := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &tmpMap)
	if err == nil {
		// TODO improve this error with guidance to the user on how to fix
		return nil, fmt.Errorf("not a map")
	}

	if len(tmpMap) != 1 {
		// TODO improve this error with guidance to the user on how to fix
		return nil, fmt.Errorf("expected exactly 1 key in map (the name of the step type), but found %d key(s): %v", len(tmpMap), lo.Keys(tmpMap))
	}

	stepName := lo.Keys(tmpMap)[0]
	return UnmarshalFromStepName(StepName(stepName), data)
}

func UnmarshalFromStepName(stepName StepName, data []byte) (Step, error) {
	unmarshalFunc, ok := unmarshalFuncsMap[stepName]
	if !ok {
		return nil, fmt.Errorf("unknown step: %s", stepName)
	}

	return unmarshalFunc(data)
}
