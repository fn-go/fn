package fnfile

import (
	"io"
	"sync"

	"github.com/hashicorp/go-multierror"
)

type ResponseWriter interface {
	Defer(StepHandler)
	OutWriter() io.Writer
	ErrOutWriter() io.Writer
	Error(error)
	ErrorOrNil() error
}

type StdResponseWriter struct {
	mu        sync.Mutex
	stdOut    io.Writer
	stdErrOut io.Writer
	err       *multierror.Error
	deferrals []StepHandler
}

func (w *StdResponseWriter) OutWriter() io.Writer {
	return w.stdOut
}

func (w *StdResponseWriter) ErrOutWriter() io.Writer {
	return w.stdErrOut
}

func (w *StdResponseWriter) Error(err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.err = multierror.Append(w.err, err)
}

func (w *StdResponseWriter) ErrorOrNil() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.err.ErrorOrNil()
}

func (w *StdResponseWriter) Defer(h StepHandler) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.deferrals = append(w.deferrals, h)
}

func NewStdResponseWriter(out, errOut io.Writer) *StdResponseWriter {
	return &StdResponseWriter{
		stdOut:    out,
		stdErrOut: errOut,
	}
}

type DiscardResponseWriter struct {
	ResponseWriter
}

func (w DiscardResponseWriter) OutWriter() io.Writer {
	return io.Discard
}

func (w DiscardResponseWriter) ErrOutWriter() io.Writer {
	return io.Discard
}

func (w DiscardResponseWriter) Error(_ error) {}

func (w DiscardResponseWriter) Defer(StepHandler) {}

func (w DiscardResponseWriter) ErrorOrNil() error { return nil }

func WithNewDeferrals(parent ResponseWriter) (ResponseWriter, *[]StepHandler) {
	d := &[]StepHandler{}

	return &deferralResponseWriter{
		ResponseWriter: parent,
		deferrals:      d,
	}, d
}

type deferralResponseWriter struct {
	ResponseWriter

	mu        sync.Mutex
	deferrals *[]StepHandler
}

func (d *deferralResponseWriter) Defer(h StepHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	*d.deferrals = append(*d.deferrals, h)
}
