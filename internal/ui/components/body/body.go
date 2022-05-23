package body

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `
Today’s Menu

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] René Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appétit!
`

type Model struct {
	viewport viewport.Model
}

func New() (Model, error) {
	vp := viewport.New(78, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(glamour.WithAutoStyle())
	if err != nil {
		return Model{}, err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return Model{}, err
	}

	vp.SetContent(str)

	return Model{
		viewport: vp,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(updateMsg tea.Msg) (tea.Model, tea.Cmd) {
	switch resolvedMsg := updateMsg.(type) {
	default:
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(resolvedMsg)
		return m, cmd
	}
}

func (m Model) View() string {
	return m.viewport.View()
}
