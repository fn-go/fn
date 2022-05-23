package keys

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
	//PageUp     key.Binding
	//PageDown   key.Binding
	PrevTab    key.Binding
	NextTab    key.Binding
	FocusLeft  key.Binding
	FocusRight key.Binding
	Help       key.Binding
	Quit       key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		//{k.PageUp, k.PageDown},
		{k.PrevTab, k.NextTab},
		{k.FocusLeft, k.FocusRight},
		{k.Help, k.Quit},
	}
}

func Keys() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys(tea.KeyUp.String(), "w"),
			key.WithHelp("↑/s", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys(tea.KeyDown.String(), "s"),
			key.WithHelp("↓/s", "move down"),
		),
		//PageUp: key.NewBinding(
		//	key.WithKeys("pgup", "b"),
		//	key.WithHelp("b/pgup", "page up/back"),
		//),
		//PageDown: key.NewBinding(
		//	key.WithKeys("pgdown", "f"),
		//	key.WithHelp("f/pgdown", "page down/forward"),
		//),
		PrevTab: key.NewBinding(
			key.WithKeys(tea.KeyLeft.String(), "a"),
			key.WithHelp("/a", "previous section"),
		),
		NextTab: key.NewBinding(
			key.WithKeys(tea.KeyRight.String(), "d"),
			key.WithHelp("/d", "next section"),
		),
		FocusLeft: key.NewBinding(
			// like vim
			// https://github.com/christoomey/vim-tmux-navigator
			key.WithKeys(tea.KeyCtrlH.String()),
			key.WithHelp("ctrl+/a", "focus previous panel"),
		),
		FocusRight: key.NewBinding(
			// like vim
			// https://github.com/christoomey/vim-tmux-navigator
			key.WithKeys(tea.KeyCtrlL.String()),
			key.WithHelp("ctrl+/d", "focus next panel"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", tea.KeyEscape.String(), tea.KeyCtrlC.String()),
			key.WithHelp("q/ctrl+c/esc", "quit"),
		),
	}
}
