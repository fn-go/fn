package main

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/oklog/run"

	"github.com/go-fn/fn/internal/clioptions/iostreams"
	"github.com/go-fn/fn/internal/cmds"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

const binaryName = "fn"

func main() {
	fmt.Fprintf(os.Stderr, "%s version: %s %s %s %s", binaryName, version, commit, date, builtBy)

	ctx, cancel := context.WithCancel(context.Background())

	ioStreams := iostreams.NewStdIOStreams()
	rootCmd := cmds.NewRootCmd(ioStreams)

	var group run.Group

	group.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))
	group.Add(func() error {
		return rootCmd.ExecuteContext(ctx)
	}, func(err error) {
		cancel()
	})

	err := group.Run()
	if err != nil {
		_, err2 := fmt.Fprintf(ioStreams.ErrOut(), "error: %s", err)
		if err2 != nil {
			panic(err2)
		}
		os.Exit(1)
	}
}
