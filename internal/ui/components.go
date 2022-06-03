package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ComponentsMap[T comparable] map[T]tea.Model
