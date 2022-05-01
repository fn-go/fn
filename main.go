package main

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/oklog/run"
	"github.com/spf13/cobra"

	"github.com/go-fn/fn/internal/clioptions/iostreams"
	"github.com/go-fn/fn/internal/cmdgroups"
	"github.com/go-fn/fn/internal/cmds"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ioStreams := iostreams.NewStdIOStreams()
	cmdGroups := cmdgroups.CommandGroups{
		{
			Message:  "Basic Commands (Beginner)",
			Commands: []*cobra.Command{},
		},
	}

	rootCmd := cmds.NewRootCmd(ioStreams)
	cmdgroups.ActsAsRootCommand(rootCmd, nil, cmdGroups...)

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
