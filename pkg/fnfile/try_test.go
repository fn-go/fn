package fnfile

import (
	"testing"
)

func TestTry_Visit(t *testing.T) {
	visitTest(t, "Try", &Try{})
}
