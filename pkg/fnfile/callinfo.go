package fnfile

import (
	"context"
	"io"
)

type CallInfo struct {
	ctx context.Context
	in  io.Reader
}

func (c *CallInfo) In() io.Reader {
	return nil
}

func (c *CallInfo) Context() context.Context {
	return c.ctx
}

func (c *CallInfo) CloneWith(ctx context.Context) *CallInfo {
	return &CallInfo{
		ctx: ctx,
		in:  c.in,
	}
}

func NewCallInfo(ctx context.Context, in io.Reader) *CallInfo {
	return &CallInfo{
		ctx: ctx,
		in:  in,
	}
}
