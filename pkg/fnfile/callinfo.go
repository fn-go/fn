package fnfile

import (
	"bytes"
	"context"
	"io"

	"github.com/ghostsquad/go-timejumper"
)

type FnContext struct {
	ctx    context.Context
	in     io.Reader
	clock  timejumper.Clock
	fnFile FnFile
}

func (c *FnContext) In() io.Reader {
	return nil
}

func (c *FnContext) Context() context.Context {
	return c.ctx
}

func (c *FnContext) FnFile() FnFile {
	return c.fnFile
}

func (c *FnContext) CloneWith(ctx context.Context) *FnContext {
	return &FnContext{
		ctx:   ctx,
		in:    c.in,
		clock: c.clock,
	}
}

type CallInfoOptions struct {
	Clock timejumper.Clock
	In    io.Reader
}

type CallInfoOption func(options *CallInfoOptions)

func NewCallInfo(ctx context.Context, options ...CallInfoOption) *FnContext {
	opts := &CallInfoOptions{
		Clock: timejumper.RealClock{},
		In:    &bytes.Buffer{},
	}

	for _, o := range options {
		o(opts)
	}

	return &FnContext{
		ctx:   ctx,
		in:    opts.In,
		clock: opts.Clock,
	}
}
