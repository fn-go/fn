package fnfile

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type unmarshalTestCase[T Step] struct {
	name    string
	data    string
	want    T
	wantErr assert.ErrorAssertionFunc
}

func newUnmarshalTestCase[T Step](name, data string, creator func(value string) T) unmarshalTestCase[T] {
	return unmarshalTestCase[T]{
		name: name,
		data: fmt.Sprintf("%q", data),
		want: creator(data),
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func testUnmarshal[T Step](t *testing.T, unmarshaller func(data []byte) (T, error), tests ...unmarshalTestCase[T]) {
	var noop T
	typeOfT := getType(noop)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshaller([]byte(tt.data))
			if tt.wantErr != nil {
				if !tt.wantErr(t, err, fmt.Sprintf("Unmarshal%s(%v)", typeOfT, tt.data)) {
					return
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equalf(t, tt.want, got, "Unmarshal%s(%v)", typeOfT, tt.data)
		})
	}
}
