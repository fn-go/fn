package fnfile

import (
	"testing"
)

func TestReturnSpec_Visit2_Visit(t *testing.T) {
	visitTest(t, "Return", &ReturnSpec{})
}
