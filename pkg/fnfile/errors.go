package fnfile

import (
	"sync"

	"github.com/hashicorp/go-multierror"
)

type ThreadSafeMultiError struct {
	mu sync.Mutex
	*multierror.Error
}

func (m *ThreadSafeMultiError) Append(errs ...error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Error = multierror.Append(m.Error, errs...)
}
