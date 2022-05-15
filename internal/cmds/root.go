package cmds

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/go-fn/fn/internal/clioptions/iostreams"
	"github.com/go-fn/fn/internal/cmdgroups"
)

const (
	// https://patorjk.com/software/taag/#p=testall&f=Patorjk's%20Cheese&t=fn
	// larry 3d
	longDescLarry3d = `
   ___           
 /'___\          
/\ \__/   ___    
\ \ ,__\/' _  \  
 \ \ \_//\ \/\ \ 
  \ \_\ \ \_\ \_\
   \/_/  \/_/\/_/
`

	// lean
	longDescLean = `
      _/_/
   _/      _/_/_/
_/_/_/_/  _/    _/
 _/      _/    _/
_/      _/    _/

`
)

func NewRootCmd(ioStreams iostreams.IOStreams) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "fn [flags] <fn>",
		Short: "fn - a simple replacement for Make\n",
		Long:  longDescLean,
		Args:  cobra.ArbitraryArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return runHelp(cmd)
			}

			fmt.Printf("root %s\n", args[0])
			return nil
		},
	}

	cmdGroups := cmdgroups.CommandGroups{
		{
			Commands: []*cobra.Command{},
		},
	}

	flagset := rootCmd.PersistentFlags()
	globalOpts := NewGlobalOptions()
	flagset.AddFlagSet(globalOpts.FlagSet())

	cmdgroups.ActsAsRootCommand(rootCmd, nil, cmdGroups...)

	rootCmd.AddCommand(NewOptionsCmd(ioStreams))

	return rootCmd
}

func runHelp(cmd *cobra.Command) error {
	return cmd.Help()
}

// TODO Learn from others: https://github.com/infrahq/infra/blob/main/internal/cmd/cmd.go
