package engine

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/go-fn/fn/pkg/fnfile"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func setupEngine(out, errOut *bytes.Buffer) *Engine {
	return New(func(opts *Options) {
		opts.Writer = fnfile.NewStdResponseWriter(out, errOut)
	})
}

func ExampleEngine_Run_do_sh() {
	out := bytes.Buffer{}
	errOut := bytes.Buffer{}

	eng := setupEngine(&out, &errOut)

	ctx := context.TODO()

	fn := fnfile.Fn{
		Name: "test",
		Do: fnfile.Do{
			Steps: fnfile.Steps{
				fnfile.Sh{
					Run: `echo "hello"`,
				},
			},
		},
	}

	panicOnError(eng.Run(ctx, fn))

	fmt.Println(out.String())
	// Output: hello
}

func ExampleEngine_Run_do_nested() {
	out := bytes.Buffer{}
	errOut := bytes.Buffer{}

	eng := setupEngine(&out, &errOut)

	ctx := context.TODO()

	fn := fnfile.Fn{
		Name: "test",
		Do: fnfile.Do{
			Steps: fnfile.Steps{
				fnfile.Sh{
					Run: `echo "one"`,
				},
				fnfile.Do{
					Steps: fnfile.Steps{
						fnfile.Sh{
							Run: `echo "nested two"`,
						},
					},
				},
				fnfile.Sh{
					Run: `echo "three"`,
				},
			},
		},
	}

	panicOnError(eng.Run(ctx, fn))

	fmt.Println(out.String())
	// Output: one
	// nested two
	// three
}

func ExampleEngine_Run_do_stop_on_failure() {
	out := bytes.Buffer{}
	errOut := bytes.Buffer{}

	eng := setupEngine(&out, &errOut)

	ctx := context.TODO()

	fn := fnfile.Fn{
		Name: "test",
		Do: fnfile.Do{
			Steps: fnfile.Steps{
				fnfile.Sh{
					Run: `echo "first"`,
				},
				fnfile.Sh{
					Run: `false`,
				},
				fnfile.Sh{
					Run: `echo "last"`,
				},
			},
		},
	}

	err := eng.Run(ctx, fn)

	fmt.Printf("result: %s, error: %s", strings.TrimSpace(out.String()), err)
	// Output: result: first, error: 1 error occurred:
	// 	* running sh [false]: exit status 1
}

func ExampleEngine_Run_defer() {
	out := bytes.Buffer{}
	errOut := bytes.Buffer{}

	eng := setupEngine(&out, &errOut)

	ctx := context.TODO()

	fn := fnfile.Fn{
		Name: "test",
		Do: fnfile.Do{
			Steps: fnfile.Steps{
				fnfile.DeferSpec{
					Do: fnfile.Do{
						Steps: fnfile.Steps{
							fnfile.Sh{
								Run: `echo "defer 1: Sub steps of a defer step..."`,
							},
							fnfile.Sh{
								Run: `echo "defer 1: run sequentially"`,
							},
						},
					},
				},
				fnfile.DeferSpec{
					Do: fnfile.Do{
						Steps: fnfile.Steps{
							fnfile.Sh{
								Run: `echo "defer 2: Deferred function calls are executed in Last In First Out order. This should show up first."`,
							},
						},
					},
				},
				fnfile.Sh{
					Run: `echo "Hello, World"`,
				},
			},
		},
	}

	panicOnError(eng.Run(ctx, fn))

	fmt.Printf(out.String())
	// Output: Hello, World
	// defer 2: Deferred function calls are executed in Last In First Out order. This should show up first.
	// defer 1: Sub steps of a defer step...
	// defer 1: run sequentially
}
