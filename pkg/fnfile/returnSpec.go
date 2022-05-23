package fnfile

// ReturnSpec is a step hook which allows a fn to return early (skipping subsequent steps) just like
// a `return` statement in Go.
type ReturnSpec struct {
	StepMeta
}

func (r *ReturnSpec) Exec(w ResponseWriter, c *CallInfo) {
	validateHandlerParams(w, c)
	c.Context().Done()
}
