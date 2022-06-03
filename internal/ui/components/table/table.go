package table

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/go-fn/fn/internal/util"
)

// Model is Table, a presentational component
// ref: https://flaviocopes.com/react-presentational-vs-container-components/
// ref: https://www.jetbrains.com/webstorm/guide/tutorials/react_typescript_tdd/presentation_components/

// Would like the table to be responsive
// https://medium.com/appnroll-publication/5-practical-solutions-to-make-responsive-data-tables-ff031c48b122
// https://www.uxmatters.com/mt/archives/2020/07/designing-mobile-tables.php
// https://css-tricks.com/responsive-data-table-roundup/

type Header interface {
	View() string
	Init() tea.Cmd
	Update(tea.Msg) (Header, tea.Cmd)
	Key() string
	MinWidth() int
}

type Headers interface {
	List() []Header

	InitHeader(string) tea.Cmd
	UpdateHeader(string, tea.Msg) (Headers, tea.Cmd)
}

type Field interface {
	View() string
	Key() string
	FilterValue() string
	MinWidth() int
}

type Fields map[string]Field

type Item interface {
	Init() tea.Cmd
	Update(tea.Msg) (Item, tea.Cmd)

	Fields() Fields
	InitField(string) tea.Cmd
	UpdateField(string, tea.Msg) (Item, tea.Cmd)
}

type Items []Item

type Model struct {
	state state
}

func New(headers Headers, items Items) Model {
	headerList := headers.List()

	headersOrder := make([]string, len(headerList))
	headersMap := make(map[string]Header, len(headerList))

	for i, h := range headerList {
		k := h.Key()

		headersOrder[i] = k
		headersMap[k] = h
	}

	return Model{
		state: state{
			items:        items,
			headersOrder: headersOrder,
			headersMap:   headersMap,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(updateMsg tea.Msg) (tea.Model, tea.Cmd) {
	switch resolvedMsg := updateMsg.(type) {
	case util.Dimensions:
		m.state.width = resolvedMsg.Width
		m.state.height = resolvedMsg.Height
	}

	return m, nil
}

func (m Model) View() string {
	// perform 2 passes on items, the first to get ideal widths of everything
	// second pass will be to tell items what width they really need to be (which will be based on max widths of sibling items)

	columnWidths := make(map[string]int, len(m.state.headersMap))

	for hk, h := range m.state.headersMap {
		columnWidths[hk] = len(h.View())
	}

	for _, i := range m.state.items {
		for fk, f := range i.Fields() {
			if _, ok := m.state.headersMap[fk]; !ok {
				continue
			}

			columnWidths[fk] = util.Max(columnWidths[fk], len(f.View()))
		}
	}

	// we still need to know what our maximum width is for the table
	// and adjust specific columns accordingly

	for hk, h := range m.state.headersMap {
		// just going to ignore commands for right now
		m.state.headersMap[hk], _ = h.Update(util.Dimensions{Width: columnWidths[hk], Height: 1})
	}

	for idx, i := range m.state.items {
		for fk := range i.Fields() {
			if _, ok := m.state.headersMap[fk]; !ok {
				continue
			}

			// just going to ignore commands for right now
			m.state.items[idx], _ = i.UpdateField(fk, util.Dimensions{Width: columnWidths[fk], Height: 1})
		}
	}

	return ""
}

var _ tea.Model = Model{}
