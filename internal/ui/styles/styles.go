package styles

import "github.com/charmbracelet/lipgloss"

func MainTextStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(DefaultTheme().MainText).
		Bold(true)
}
