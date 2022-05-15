package fnfile

// ReturnSpec is a step hook which allows a fn to return early (skipping subsequent steps) just like
// a `return` statement in Go.
type ReturnSpec struct {
	*StepCommon
}

func (r *ReturnSpec) Exec(_ ResponseWriter, c *CallInfo) {
	c.Context().Done()
}
