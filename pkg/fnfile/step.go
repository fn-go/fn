package fnfile

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/samber/lo"
)

type Steps []Step

func (steps *Steps) UnmarshalJSON(data []byte) (err error) {
	var tmpList []json.RawMessage
	err = json.Unmarshal(data, &tmpList)
	if err != nil {
		return fmt.Errorf("unmarshalling steps to list: %w", err)
	}

	var mErr *multierror.Error

	tmpSteps := make([]Step, len(tmpList))
	for i, s := range tmpList {
		tmpSteps[i], err = UnmarshalStep(s)
		mErr = multierror.Append(mErr, err)
	}

	*steps = tmpSteps

	err = mErr.ErrorOrNil()
	return
}

type Step interface {
	Accept(visitor StepVisitor)
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

func NewStepVisitor(w ResponseWriter, c *FnContext) StepVisitor {
	return StepVisitor{
		VisitDefer:    HandleStepWith[DeferSpec](w, c),
		VisitDo:       HandleStepWith[Do](w, c),
		VisitFnStep:   HandleStepWith[FnStepSpec](w, c),
		VisitMatrix:   HandleStepWith[Matrix](w, c),
		VisitParallel: HandleStepWith[Parallel](w, c),
		VisitReturn:   HandleStepWith[ReturnSpec](w, c),
		VisitSh:       HandleStepWith[Sh](w, c),
	}
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
)

func UnmarshalStep(data []byte) (Step, error) {
	tmpMap := make(map[string]json.RawMessage)

	var err error
	err = json.Unmarshal(data, &tmpMap)
	if err != nil {
		// TODO improve this error with guidance to the user on how to fix
		return nil, fmt.Errorf("ambiguous step, no name key")
	}

	if len(tmpMap) == 0 {
		// TODO improve this error with guidance to the user on how to fix
		return nil, fmt.Errorf("empty step? not sure what to do")
	}

	if len(tmpMap) > 1 {
		// TODO improve this error with guidance to the user on how to fix
		return nil, fmt.Errorf("too many keys, expected only 1, the name of the step type")
	}

	stepName := lo.Keys(tmpMap)[0]

	switch StepName(stepName) {
	case DeferStep:
		return UnmarshalDefer(data)
	case DoStep:
		return UnmarshalToDo(data)
	case FnStep:
		return UnmarshalFnStep(data)
	case MatrixStep:
		return UnmarshalMatrix(data)
	case ParallelStep:
		return UnmarshalParallel(data)
	case ReturnStep:
		return UnmarshalReturn(data)
	case ShStep:
		return UnmarshalSh(data)
	default:
		panic(fmt.Errorf("unknown step: %s", stepName))
	}
}

func HandleStepWith[T StepHandler](w ResponseWriter, c *FnContext) func(h T) {
	return func(h T) {
		h.Handle(w, c)
	}
}
