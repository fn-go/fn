package fnfile

import (
	"testing"
)

func TestUnmarshalDefer(t *testing.T) {
	testUnmarshal(t, UnmarshalDefer,
		newUnmarshalTestCase(
			"string shorthand",
			`echo "hello"`,
			func(value string) DeferSpec {
				return DeferSpec{
					Do: Do{
						Steps: Steps{
							Sh{
								Run: value,
							},
						},
					},
				}
			},
		),
		unmarshalTestCase[DeferSpec]{
			name: "array shorthand",
			data: `
[
  "echo \"first\"",
  "echo \"second\""
]
`,
			want: DeferSpec{
				Do: Do{
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
		},
	)
}
