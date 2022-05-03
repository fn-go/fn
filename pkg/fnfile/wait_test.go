package fnfile

import (
	"testing"
)

func TestWait_Visit(t *testing.T) {
	visitTest(t, "Wait", &Wait{})
}
