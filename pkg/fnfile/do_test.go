package fnfile

import (
	"testing"
)

func TestDo_Visit(t *testing.T) {
	visitTest(t, "Do", &Do{})
}
