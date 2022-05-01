package variable

import (
	"github.com/hashicorp/go-multierror"
)

// Cyclical is an error indicating a cycle of dependencies
type Cyclical interface {
	error
	Cyclical() bool
}

type cyclicalError struct {
	// Message is the Human-readable message.
	Message string `json:"message"`
	// In is the var which caught the cycle
	In string `json:"in"`
	// From is the caller/invoker of the var
	From string `json:"from"`

	// Err is the nested error.
	Cause error `json:"-"`
}

func (e cyclicalError) Error() string {
	return e.Message
}

func (e cyclicalError) Cyclical() bool {
	return true
}

func (e cyclicalError) Unwrap() error {
	return e.Cause
}

// NewCyclicalError returns a struct adhering to the Cyclical interface.
//goland:noinspection GoExportedFuncWithUnexportedType
func NewCyclicalError(in, from string, causes ...error) cyclicalError {
	var mErr *multierror.Error

	if len(causes) > 1 {
		for _, c := range causes {
			mErr = multierror.Append(mErr, c)
		}
	}

	return cyclicalError{
		Message: "cyclical request detected for lazy var: [ " + in + " ]: from: [ " + from + " ]",
		In:      in,
		From:    from,
		Cause:   mErr.ErrorOrNil(),
	}
}

// IsCyclical returns true if err is Cyclical.
func IsCyclical(err error) bool {
	te, ok := err.(Cyclical)
	return ok && te.Cyclical()
}
