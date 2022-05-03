package fnfile

import (
	"testing"
)

func TestParallel_Visit(t *testing.T) {
	visitTest(t, "Parallel", &Parallel{})
}
