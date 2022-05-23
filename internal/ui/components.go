package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Viewable interface {
	View() string
}

type Updatable interface {
	Update(tea.Msg) (tea.Model, tea.Cmd)
}

type ViewableUpdatable interface {
	Viewable
	Updatable
}

type ComponentsMap[T comparable] map[T]ViewableUpdatable
