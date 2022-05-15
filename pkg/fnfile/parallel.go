package fnfile

type Parallel struct {
	*StepCommon

	Steps    []Step `json:"steps"`
	FailFast bool   `json:"failFast"`
	Limit    int    `json:"limit"`
}

// TODO implement fail fast
func (p *Parallel) Exec(w ResponseWriter, c *CallInfo) {
	for _, s := range p.Steps {
		go s.Exec(w, c)
	}
}
