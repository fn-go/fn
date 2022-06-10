package fnfile

import (
	"testing"
)

func TestUnmarshalParallel(t *testing.T) {
	testUnmarshal(t, UnmarshalParallel,
		newUnmarshalTestCase(
			"string shorthand",
			`echo "hello"`,
			func(value string) Parallel {
				return Parallel{
					Steps: Steps{
						Sh{
							Run: `echo "hello"`,
						},
					},
				}
			},
		),
		unmarshalTestCase[Parallel]{
			name: "array shorthand",
			data: `
[
  "echo \"hello\"",
  "echo \"world\""
]
`,
			want: Parallel{
				Steps: Steps{
					Sh{
						Run: `echo "hello"`,
					},
					Sh{
						Run: `echo "world"`,
					},
				},
			},
		},
	)
}
