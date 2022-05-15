package protocol

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/yaml"

	"github.com/go-fn/fn/pkg/fnfile"
)

func quickParse(data string) (file fnfile.FnFile, err error) {
	err = yaml.Unmarshal([]byte(data), &file)
	return
}

func TestYaml(t *testing.T) {
	type testcase struct {
		name    string
		example string
		want    string
		err     error
	}

	for _, tt := range []testcase{
		{
			name: "simple",
			example: `
version: '0.1'
fns:
  test: echo "hello"
`,
			want: "hello",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			f, err := quickParse(tt.example)
			require.NoError(t, err)

			out := bytes.Buffer{}
			errOut := bytes.Buffer{}

			eng := setupEngine(&out, &errOut)
			err = eng.Run(context.TODO(), f.Fns["test"])
			require.NoError(t, err)

			assert.Equal(t, tt.want, out.String())
		})
	}
}
