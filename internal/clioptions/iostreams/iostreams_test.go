package iostreams

import (
	"bytes"
	"io"
	"io/ioutil"
)

type stateful struct {
	in     io.Reader
	out    io.Writer
	errOut io.Writer
}

func (s *stateful) In() io.Reader {
	return s.in
}

func (s *stateful) Out() io.Writer {
	return s.out
}

func (s *stateful) ErrOut() io.Writer {
	return s.errOut
}

// NewStateful allows you to create your own IOStreams
func NewStateful(in io.Reader, out, errOut io.Writer) IOStreams {
	return &stateful{
		in:     in,
		out:    out,
		errOut: errOut,
	}
}

// NewDiscard returns a valid IOStreams that just discards everything
func NewDiscard() IOStreams {
	return &stateful{
		in:     &bytes.Buffer{},
		out:    ioutil.Discard,
		errOut: ioutil.Discard,
	}
}
