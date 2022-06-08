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
							Run: value,
						},
					},
				}
			},
		),
		unmarshalTestCase[Parallel]{
			name: "array shorthand",
			data: `
[
  "echo \"first\"",
  "echo \"second\""
]
`,
			want: Parallel{
				Steps: Steps{
					Sh{
						Run: `echo "first"`,
					},
					Sh{
						Run: `echo "second"`,
					},
				},
			},
		},
	)
}
