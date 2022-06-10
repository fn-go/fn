package fnfile

import (
	"fmt"
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

type GivenUp struct {
	attempts *multierror.Error
}

func (g GivenUp) Error() string {
	return fmt.Errorf("giving up: %w", g.attempts).Error()
}

func GivingUp(attempts *multierror.Error) error {
	return GivenUp{
		attempts: attempts,
	}
}
