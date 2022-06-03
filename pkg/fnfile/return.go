package fnfile

// ReturnSpec is a step hook which allows a fn to return early (skipping subsequent steps) just like
// a `return` statement in Go.
type ReturnSpec struct {
	StepMeta
}

func (r ReturnSpec) Accept(visitor StepVisitor) {
	visitor.VisitReturn(r)
}

func (r ReturnSpec) Handle(_ ResponseWriter, c *FnContext) {
	c.Context().Done()
}

func UnmarshalReturn(data []byte) (ReturnSpec, error) {
	return ReturnSpec{}, nil
}
