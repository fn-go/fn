package fnfile

type FnStepSpec struct {
	StepMeta
	Fn string `json:"fn,omitempty"`
}

func (fn FnStepSpec) Accept(visitor StepVisitor) {
	visitor.VisitFnStep(fn)
}

func (fn FnStepSpec) Handle(w ResponseWriter, c *FnContext) {
	fnfile := c.FnFile()
	if target, ok := fnfile.Fns[fn.Fn]; ok {
		target.Do.Handle(w, c)
	}
}

func UnmarshalFnStep(data []byte) (FnStepSpec, error) {
	return FnStepSpec{}, nil
}
