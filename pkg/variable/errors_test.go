package variable

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

var faker = gofakeit.New(0)

func Test_cyclicalError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  cyclicalError
		want string
	}{
		{
			name: "human readable error string",
			err:  NewCyclicalError("hello", "world"),
			want: "cyclical request detected for lazy var: hello: from: world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cyclicalError_As_CyclicalError(t *testing.T) {

}
