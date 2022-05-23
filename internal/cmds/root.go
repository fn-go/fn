package cmds

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/go-fn/fn/internal/clioptions"
	"github.com/go-fn/fn/internal/clioptions/iostreams"
	"github.com/go-fn/fn/internal/ui/app"
)

const (
	// https://patorjk.com/software/taag/#p=testall&f=Patorjk's%20Cheese&t=fn
	// larry 3d
	//	longDescLarry3d = `
	//   ___
	// /'___\
	///\ \__/   ___
	//\ \ ,__\/' _  \
	// \ \ \_//\ \/\ \
	//  \ \_\ \ \_\ \_\
	//   \/_/  \/_/\/_/
	//`

	// lean
	longDescLean = `
      _/_/
   _/      _/_/_/
_/_/_/_/  _/    _/
 _/      _/    _/
_/      _/    _/
`
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected

	tabs    struct{}
	sidebar struct{}
}

func NewRootCmd(ioStreams iostreams.IOStreams) *cobra.Command {
	globalOpts := clioptions.GlobalOptions{
		IOStreams: ioStreams,
	}

	rootCmd := &cobra.Command{
		Use:   "fn [flags] FN",
		Short: "fn - a function-oriented interpretation of Make",
		Long:  longDescLean,
		Args:  cobra.ArbitraryArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ParseFlags(args); err != nil {
				return err
			}

			if globalOpts.Interactive {
				instance, err := app.New()
				if err != nil {
					fmt.Println(fmt.Errorf("creating app: %w", err))
				}

				p := tea.NewProgram(instance, tea.WithAltScreen())
				if err := p.Start(); err != nil {
					fmt.Println(fmt.Errorf("starting program: %w", err))
					os.Exit(1)
				}
			} else {
				if len(args) == 0 {
					return runHelp(cmd)
				}
			}

			return nil
		},
	}

	flags := rootCmd.PersistentFlags()
	flags.BoolVarP(&globalOpts.Interactive, "interactive", "i", false, "interactive mode")

	rootCmd.SetHelpTemplate(`{{.Long}}
{{.Short}}

{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`)

	return rootCmd
}

func runHelp(cmd *cobra.Command) error {
	return cmd.Help()
}

// TODO Learn from others: https://github.com/infrahq/infra/blob/main/internal/cmd/cmd.go
