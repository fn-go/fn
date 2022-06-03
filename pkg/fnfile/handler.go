package fnfile

type StepHandler interface {
	Handle(w ResponseWriter, c *FnContext)
}

type StepHandlerFunc func(w ResponseWriter, c *FnContext)

func (f StepHandlerFunc) Handle(w ResponseWriter, c *FnContext) {
	f(w, c)
}
