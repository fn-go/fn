package iostreams

import (
	"io"
	"os"
)

type std struct{}

func (s *std) In() io.Reader {
	return os.Stdin
}

func (s *std) Out() io.Writer {
	return os.Stdout
}

func (s *std) ErrOut() io.Writer {
	return os.Stderr
}

func NewStdIOStreams() *std {
	return &std{}
}

type IOStreams interface {
	In() io.Reader
	Out() io.Writer
	ErrOut() io.Writer
}
