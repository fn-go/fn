package fnfile

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"sync"

	"github.com/hashicorp/go-multierror"
)

type Deferer interface {
	Defer(Handler)
}

type ResponseWriter interface {
	Deferer

	OutWriter() io.Writer
	ErrOutWriter() io.Writer
	Error(error)
	ErrorOrNil() error
}

type DiscardResponseWriter struct {
	ResponseWriter
}

func (w DiscardResponseWriter) OutWriter() io.Writer {
	return ioutil.Discard
}

func (w DiscardResponseWriter) ErrOutWriter() io.Writer {
	return ioutil.Discard
}

func (w DiscardResponseWriter) Error(_ error) {}

func (w DiscardResponseWriter) Defer(Handler) {}

func (w DiscardResponseWriter) ErrorOrNil() error { return nil }

type BufferResponseWriter struct {
	mu        sync.Mutex
	bufOut    *bytes.Buffer
	bufErrOut *bytes.Buffer
	err       *multierror.Error
	deferrals []Handler
}

func (w *BufferResponseWriter) OutWriter() io.Writer {
	return w.bufOut
}

func (w *BufferResponseWriter) ErrOutWriter() io.Writer {
	return w.bufErrOut
}

func (w *BufferResponseWriter) Error(err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.err = multierror.Append(w.err, err)
}

func (w *BufferResponseWriter) ErrorOrNil() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.err.ErrorOrNil()
}

func (w *BufferResponseWriter) Defer(h Handler) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.deferrals = append(w.deferrals, h)
}

func NewBufferResponseWriter(out, errOut *bytes.Buffer) *BufferResponseWriter {
	return &BufferResponseWriter{
		bufOut:    out,
		bufErrOut: errOut,
	}
}

func withNewDeferrals(parent ResponseWriter) (ResponseWriter, *[]Handler) {
	var d []Handler

	return &deferResponseWriter{
		ResponseWriter: parent,
		deferrals:      d,
	}, &d
}

type deferResponseWriter struct {
	ResponseWriter

	mu        sync.Mutex
	deferrals []Handler
}

func (d *deferResponseWriter) Defer(h Handler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	fmt.Println("deferer got defer")
	d.deferrals = append(d.deferrals, h)
}
