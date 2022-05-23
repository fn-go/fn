package app

type Window struct {
	Width  int
	Height int
}

type state struct {
	Window

	currentFocus      componentKey
	currentFocusIndex int
	focusable         []componentKey
}

func newState() state {
	return state{
		currentFocus:      bodyComponentKey,
		currentFocusIndex: 0,
		focusable: []componentKey{
			bodyComponentKey,
			sidebarComponentKey,
		},
	}
}
