package cmds

import (
	"github.com/spf13/cobra"

	"github.com/go-fn/fn/internal/clioptions/iostreams"
)

func NewRootCmd(ioStreams iostreams.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use: "fn",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runHelp(cmd)
		},
	}

	return cmd
}

func runHelp(cmd *cobra.Command) error {
	return cmd.Help()
}

// TODO Learn from others: https://github.com/infrahq/infra/blob/main/internal/cmd/cmd.go
