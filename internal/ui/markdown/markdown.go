package markdown

import "github.com/charmbracelet/glamour"

func MustGetMarkdownRenderer(width int) glamour.TermRenderer {
	markdownRenderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(CustomDarkStyleConfig()),
		glamour.WithWordWrap(width),
	)

	if err != nil {
		panic(err)
	}

	return *markdownRenderer
}
