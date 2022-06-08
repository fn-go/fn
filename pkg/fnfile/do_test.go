package fnfile

import (
	"testing"
)

func TestUnmarshalDo(t *testing.T) {
	testUnmarshal(t, UnmarshalDo,
		newUnmarshalTestCase(
			"string shorthand",
			`echo "hello"`,
			func(value string) Do {
				return Do{
					Steps: Steps{
						Sh{
							Run: value,
						},
					},
				}
			},
		),
		unmarshalTestCase[Do]{
			name: "array shorthand",
			data: `
[
  "echo \"first\"",
  "echo \"second\""
]
`,
			want: Do{
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
