package fnfile

import (
	"testing"
)

func TestDeferSpec_Visit(t *testing.T) {
	visitTest(t, "Defer", &DeferSpec{})
}
