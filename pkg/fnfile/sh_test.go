package fnfile

import (
	"testing"
)

func TestSh_Visit(t *testing.T) {
	visitTest(t, "Sh", &Sh{})
}
