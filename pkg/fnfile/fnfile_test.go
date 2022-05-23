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
  test: echo "hello"
`
	fnFile := fnFileFromYaml(example)

	expected := FnFile{
		Version: "0.1",
		Fns: map[string]Fn{
			"test": NewFn("test", func(fn *Fn) {
				fn.Do = NewDo(func(do *Do) {
					do.parent = *fn
					do.Steps = Steps{
						NewSh(`echo "hello"`, func(sh *Sh) {
							sh.parent = *do
						}),
					}
				})
			}),
		},
	}

	assert.Equal(t, expected, fnFile)
}
