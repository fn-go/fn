package msg

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Focus struct {
	tea.Msg
}

type Defocus struct {
	tea.Msg
}
