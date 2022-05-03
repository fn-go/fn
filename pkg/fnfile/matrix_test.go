package fnfile

import (
	"testing"
)

func TestMatrix_Visit(t *testing.T) {
	visitTest(t, "Matrix", &Matrix{})
}
