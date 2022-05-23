package fnfile

const (
	DeferSpecStepType StepType = "defer"
)

type DeferSpec struct {
	StepMeta
	Do
}

func (d DeferSpec) Exec(w ResponseWriter, c *CallInfo) {
	validateHandlerParams(w, c)
	w.Defer(d.Do)
}
