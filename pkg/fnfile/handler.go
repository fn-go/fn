package fnfile

type Handler interface {
	Exec(ResponseWriter, *CallInfo)
}

type HandlerFunc func(ResponseWriter, *CallInfo)

// Exec calls f(ctx).
func (f HandlerFunc) Exec(w ResponseWriter, c *CallInfo) {
	f(w, c)
}

func validateHandlerParams(w ResponseWriter, c *CallInfo) {
	if w == nil {
		panic("invalid ResponseWriter")
	}

	if c == nil {
		panic("invalid CallInfo")
	}
}
