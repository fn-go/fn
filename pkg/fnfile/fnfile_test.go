package fnfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func fnFileFromYaml(data string) FnFile {
	fnFile := FnFile{}
	if err := yaml.Unmarshal([]byte(data), &fnFile); err != nil {
		panic(err)
	}
	return fnFile
}

func TestFnFile_FromYaml(t *testing.T) {
	var example = `
version: '0.1'
fns:
  test:
    do: echo "hello"
`
	fnFile := fnFileFromYaml(example)

	expected := FnFile{
		Version: "0.1",
		Fns: map[string]FnDef{
			"test": {
				Name: "test",
				Do: Do{
					Steps: Steps{
						Sh{
							Run: `echo "hello"`,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, fnFile)
}
