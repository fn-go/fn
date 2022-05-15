package fnfile

import (
	"fmt"
	"sync"

	"github.com/hashicorp/go-multierror"
)

func NewThreadSafeMultiError() *threadSafeMultiError {
	return &threadSafeMultiError{}
}

type threadSafeMultiError struct {
	sync.Mutex
	*multierror.Error
}

func (m *threadSafeMultiError) Append(errs ...error) {
	m.Lock()
	defer m.Unlock()

	m.Error = multierror.Append(m.Error, errs...)
}

func NewCancelledStepError(name string, cause error) error {
	return fmt.Errorf("step cancelled: %s: %w", name, cause)
}
