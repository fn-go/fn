package msg

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/go-fn/fn/internal/util"
)

func DimensionsMsg(d util.Dimensions) tea.Msg {
	return d
}
