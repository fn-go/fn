package cmds

import (
	"fmt"
	stdos "os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hack-pad/hackpadfs/os"
	"github.com/spf13/cobra"

	"github.com/go-fn/fn/internal/clioptions"
	"github.com/go-fn/fn/internal/clioptions/iostreams"
	"github.com/go-fn/fn/internal/ui/app"
	"github.com/go-fn/fn/pkg/engine"
	"github.com/go-fn/fn/pkg/fnfile"
)

const (
	// https://patorjk.com/software/taag/#p=testall&f=Patorjk's%20Cheese&t=fn
	// larry 3d
	//	longDescLarry3d = `
	//    ___
	//  /'___\
	// /\ \__/   ___
	// \ \ ,__\/' _  \
	//  \ \ \_//\ \/\ \
	//   \ \_\ \ \_\ \_\
	//    \/_/  \/_/\/_/
	// `

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
	globalOpts := clioptions.GlobalOptions{
		IOStreams: ioStreams,
	}

	rootCmd := &cobra.Command{
		Use:           "fn [flags] FN",
		Short:         "fn - a function-oriented interpretation of Make",
		Long:          longDescLean,
		Args:          cobra.ArbitraryArgs,
		SilenceErrors: true,
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
					return fmt.Errorf("starting program: %w", err)
				}
				return nil
			}

			if len(args) == 0 {
				return runHelp(cmd)
			}

			wd, err := stdos.Getwd()
			if err != nil {
				return err
			}

			fs, err := os.NewFS().Sub(strings.TrimLeft(wd, "/"))
			if err != nil {
				return err
			}

			fnFile, err := engine.YamlFileReader(func(opts *engine.YamlFileReaderOptions) {
				opts.FS = fs.(*os.FS)
			})()
			if err != nil {
				return fmt.Errorf("getting default fnfile for reading: %w", err)
			}

			// look for
			eng := engine.New(func(opts *engine.Options) {
				opts.Writer = fnfile.NewStdResponseWriter(ioStreams.Out(), ioStreams.ErrOut())
			})

			fname := args[0]

			err = eng.Run(cmd.Context(), fnFile.Fns[fname])
			if err != nil {
				return fmt.Errorf("running fn: %s: %w", fname, err)
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
