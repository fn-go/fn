package fnfile

type DeferSpec struct {
	StepCommon
	Do
}

type DeferSpecOptions struct {
	Do                Do
	StepCommonOptions []StepCommonOption
}

type DeferSpecOption func(options *DeferSpecOptions)

func NewDeferSpec(name string, options ...DeferSpecOption) DeferSpec {
	opts := &DeferSpecOptions{}
	for _, o := range options {
		o(opts)
	}

	return DeferSpec{
		StepCommon: NewStepCommon(name, opts.StepCommonOptions...),
		Do:         opts.Do,
	}
}

func (d DeferSpec) Exec(w ResponseWriter, c *CallInfo) {
	validateHandlerParams(w, c)
	w.Defer(d.Do)
}
