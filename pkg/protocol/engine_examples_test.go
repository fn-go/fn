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

	fn := fnfile.NewFn("test", func(fn *fnfile.Fn) {
		fn.Do = fnfile.NewDo(func(do *fnfile.Do) {
			do.Steps = fnfile.Steps{
				fnfile.NewSh(`echo "hello"`),
			}
		})
	})

	panicOnError(eng.Run(ctx, fn))

	fmt.Println(out.String())
	// Output: hello
}

func ExampleEngine_Run_do_nested() {
	out := bytes.Buffer{}
	errOut := bytes.Buffer{}

	eng := setupEngine(&out, &errOut)

	ctx := context.TODO()

	fn := fnfile.NewFn("test", func(fn *fnfile.Fn) {
		fn.Do = fnfile.NewDo(func(do *fnfile.Do) {
			do.Steps = fnfile.Steps{
				fnfile.NewSh(`echo "hello from above"`),
				fnfile.NewDo(func(do *fnfile.Do) {
					do.Steps = fnfile.Steps{
						fnfile.NewSh(`echo "hello, I'm nested!"`),
					}
				}),
				fnfile.NewSh(`echo "goodbye from above"`),
			}
		})
	})

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

	fn := fnfile.NewFn("test", func(fn *fnfile.Fn) {
		fn.Do = fnfile.NewDo(func(do *fnfile.Do) {
			do.Steps = fnfile.Steps{
				fnfile.NewSh(`echo "first"`),
				fnfile.NewSh(`false`),
				fnfile.NewSh(`echo "last"`),
			}
		})
	})

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

	fn := fnfile.NewFn("test", func(fn *fnfile.Fn) {
		fn.Do = fnfile.NewDo(func(do *fnfile.Do) {
			do.Steps = fnfile.Steps{
				fnfile.NewDeferSpec(func(deferSpec *fnfile.DeferSpec) {
					deferSpec.Name = "1defer"
					deferSpec.Steps = fnfile.Steps{
						fnfile.NewSh(`echo "defer 1: Sub steps of a defer step..."`),
						fnfile.NewSh(`echo "defer 1: run sequentially"`),
					}
				}),
				fnfile.NewDeferSpec(func(deferSpec *fnfile.DeferSpec) {
					deferSpec.Name = "defer 2"
					deferSpec.Steps = fnfile.Steps{
						fnfile.NewSh(`echo "defer 2: Deferred function calls are executed in Last In First Out order. This should show up first."`),
					}
				}),
				fnfile.NewSh(`echo "Hello, World"`),
			}
		})
	})

	panicOnError(eng.Run(ctx, fn))

	fmt.Printf(out.String())
	// Output: Hello, World
	// defer 2: Deferred function calls are executed in Last In First Out order. This should show up first.
	// defer 1: Sub steps of a defer step...
	// run sequentially
}
