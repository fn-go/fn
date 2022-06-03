package listviewport

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/go-fn/fn/internal/util"
)

type HeightAwareModel interface {
	tea.Model
	Height() int
}

// Model does not currently support a dynamically changing list
type Model struct {
	viewport viewport.Model

	state state
}

type Options struct {
	util.Dimensions
}

type Option func(options *Options)

func NewModel(items []HeightAwareModel, options ...Option) (Model, error) {
	opts := &Options{}

	for _, o := range options {
		o(opts)
	}

	model := Model{
		state: state{
			items: items,
		},
		viewport: viewport.Model{
			Width:  opts.Dimensions.Width,
			Height: opts.Dimensions.Height,
		},
	}

	return model, nil
}
