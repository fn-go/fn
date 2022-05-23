package fnfile

type FnStep struct {
	StepMeta
	Fn string `json:"fn,omitempty"`
}

func (d FnStep) Exec(w ResponseWriter, c *CallInfo) {
	validateHandlerParams(w, c)
}
