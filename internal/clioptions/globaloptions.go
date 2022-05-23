package clioptions

import (
	"github.com/go-fn/fn/internal/clioptions/iostreams"
)

type GlobalOptions struct {
	Interactive bool
	IOStreams   iostreams.IOStreams
}
