package protocol

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
	writer := fnfile.NewBufferResponseWriter(out, errOut)

	eng, err := New(func(opts *EngineOptions) {
		opts.Writer = writer
	})
	panicOnError(err)

	return eng
}

func ExampleEngine_Run_do_sh() {
	out := bytes.Buffer{}
	errOut := bytes.Buffer{}

	eng := setupEngine(&out, &errOut)

	ctx := context.TODO()

	fn := fnfile.Fn{
		Do: fnfile.Steps{
			fnfile.NewSh("hello", `echo "hello"`),
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
		Do: fnfile.Steps{
			fnfile.NewSh("1", `echo "hello from above"`),
			fnfile.NewDo("2", func(options *fnfile.DoOptions) {
				options.Steps = fnfile.Steps{
					fnfile.NewSh("3", `echo "hello, I'm nested!"`),
				}
			}),
			fnfile.NewSh("4", `echo "goodbye from above"`),
		},
	}

	panicOnError(eng.Run(ctx, fn))

	fmt.Println(out.String())
	// Output: hello from above
	// hello, I'm nested!
	// goodbye from above
}

func ExampleEngine_Run_do_stop_on_failure() {
	out := bytes.Buffer{}
	errOut := bytes.Buffer{}

	eng := setupEngine(&out, &errOut)

	ctx := context.TODO()

	fn := fnfile.Fn{
		Do: fnfile.Steps{
			fnfile.NewSh("1", `echo "first"`),
			fnfile.NewSh("fail", "false"),
			fnfile.NewSh("3", `echo "last"`),
		},
	}

	err := eng.Run(ctx, fn)

	fmt.Printf("result: %s, error: %s", strings.TrimSpace(out.String()), err)
	// Output: result: first, error: 1 error occurred:
	// 	* running sh [fail]: exit status 1
}

func ExampleEngine_Run_defer() {
	out := bytes.Buffer{}
	errOut := bytes.Buffer{}

	eng := setupEngine(&out, &errOut)

	ctx := context.TODO()

	fn := fnfile.Fn{
		Do: fnfile.Steps{
			fnfile.NewDeferSpec("1defer", func(options *fnfile.DeferSpecOptions) {
				options.Do = fnfile.NewDo("1defer.do", func(options *fnfile.DoOptions) {
					options.Steps = fnfile.Steps{
						fnfile.NewSh("defer1.do.sh", `echo "defer 1: Sub steps of a defer step..."`),
						fnfile.NewSh("defer1.do.sh", `echo "run sequentially"`),
					}
				})
			}),
			fnfile.NewDeferSpec("2defer", func(options *fnfile.DeferSpecOptions) {
				options.Do = fnfile.NewDo("2defer.do", func(options *fnfile.DoOptions) {
					options.Steps = fnfile.Steps{
						fnfile.NewSh("defer2.do.sh", `echo "defer 2: Deferred function calls are executed in Last In First Out order. This should show up first."`),
					}
				})
			}),
			fnfile.NewSh("3sh", `echo "Hello, World"`),
		},
	}

	panicOnError(eng.Run(ctx, fn))

	fmt.Printf(out.String())
	// Output: Hello, World
	// defer 2: Deferred function calls are executed in Last In First Out order. This should show up first.
	// defer 1: Sub steps of a defer step...
	// run sequentially
}
