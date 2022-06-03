package app

import (
	"github.com/go-fn/fn/internal/util"
)

type state struct {
	windowDimensions util.Dimensions

	headerDimensions  util.Dimensions
	bodyDimensions    util.Dimensions
	sidebarDimensions util.Dimensions
	footerDimensions  util.Dimensions

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
