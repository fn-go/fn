package app

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/go-fn/fn/internal/ui"
	"github.com/go-fn/fn/internal/ui/components/body"
	"github.com/go-fn/fn/internal/ui/components/footer"
	"github.com/go-fn/fn/internal/ui/components/sidebar"
	"github.com/go-fn/fn/internal/ui/components/tabs"
	"github.com/go-fn/fn/internal/ui/keys"
	"github.com/go-fn/fn/internal/ui/msg"
)

type componentKey string

const (
	tabsComponentKey    componentKey = "tabs"
	bodyComponentKey    componentKey = "body"
	sidebarComponentKey componentKey = "sidebar"
	footerComponentKey  componentKey = "footer"
)

type Model struct {
	keys  keys.KeyMap
	err   error
	state state

	components ui.ComponentsMap[componentKey]
}

func New() (Model, error) {
	bodyComp, err := body.New()
	if err != nil {
		return Model{}, err
	}

	sidebarComp, err := sidebar.New()
	if err != nil {
		return Model{}, err
	}

	s := newState()
	s.sidebarDimensions.Width = 50

	return Model{
		keys:  keys.Keys(),
		state: s,

		components: ui.ComponentsMap[componentKey]{
			tabsComponentKey:    tabs.New(),
			bodyComponentKey:    bodyComp,
			sidebarComponentKey: sidebarComp,
			footerComponentKey:  footer.New(),
		},
	}, nil
}

func (m Model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, len(m.components))

	i := 0
	for k, c := range m.components {
		cmd := c.Init()
		m.components[k], cmds[i] = c.Update(cmd)
	}

	return tea.Batch(cmds...)
}

func (m Model) Update(updateMsg tea.Msg) (tea.Model, tea.Cmd) {
	switch resolvedMsg := updateMsg.(type) {
	case tea.KeyMsg:
		return m.onKeyPress(resolvedMsg)

	case tea.WindowSizeMsg:
		return m.onWindowSizeChange(resolvedMsg)
	}

	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString(m.components[tabsComponentKey].View())
	s.WriteString("\n")
	s.WriteString(lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.components[bodyComponentKey].View(),
		m.components[sidebarComponentKey].View(),
	))
	s.WriteString("\n")
	s.WriteString(m.components[footerComponentKey].View())
	return s.String()
}

func (m Model) onKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.PrevTab):

	case key.Matches(msg, m.keys.NextTab):

	case key.Matches(msg, m.keys.FocusLeft):
		return m.refocus(-1)

	case key.Matches(msg, m.keys.FocusRight):
		return m.refocus(1)

	case key.Matches(msg, m.keys.Quit):
		return m, tea.Quit
	default:
		return m.updateComponent(m.state.currentFocus, msg)
	}

	return m, nil
}

//goland:noinspection GoAssignmentToReceiver
func (m Model) refocus(change int) (tea.Model, tea.Cmd) {
	requestedIndex := m.state.currentFocusIndex + change
	if requestedIndex < 0 {
		return m, nil
	}
	if requestedIndex > len(m.state.focusable)-1 {
		return m, nil
	}

	cmds := make([]tea.Cmd, 2)

	m, cmds[0] = m.updateComponent(m.state.currentFocus, msg.Defocus{})

	m.state.currentFocusIndex = m.state.currentFocusIndex + change

	m, cmds[1] = m.updateComponent(m.state.currentFocus, msg.Focus{})
	m.state.currentFocus = m.state.focusable[m.state.currentFocusIndex]

	return m, tea.Batch(cmds...)
}

func (m Model) updateComponent(key componentKey, teaMsg tea.Msg) (Model, tea.Cmd) {
	compUpdate, compCmd := m.components[key].Update(teaMsg)
	m.components[key] = compUpdate
	return m, compCmd
}

func (m Model) onWindowSizeChange(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.state.windowDimensions.Height = msg.Height
	m.state.windowDimensions.Width = msg.Width

	m.state.bodyDimensions.Width = m.state.windowDimensions.Width - m.state.sidebarDimensions.Width

	cmds := make([]tea.Cmd, len(m.components))

	i := 0
	for compType, comp := range m.components {
		newModel, newCmd := comp.Update(msg)
		cmds[i] = newCmd
		m.components[compType] = newModel

		i++
	}

	return m, tea.Batch(cmds...)
}
