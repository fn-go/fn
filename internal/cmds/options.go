// TODO Learn from others: https://github.com/infrahq/infra/blob/main/internal/cmd/cmd.go
// TODO Learn from others: https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/options/options.go

package cmds

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"

	// Note that this is being replaced in go.mod with the cornfeedhobo fork
	"github.com/spf13/pflag"

	"github.com/go-fn/fn/internal/clioptions/iostreams"
	"github.com/go-fn/fn/internal/cmdgroups"
)

func NewOptionsCmd(ioStreams iostreams.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fn-options",
		Short: "Print the list of flags inherited by all commands",
		Long:  "Print the list of flags inherited by all commands",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("options")
			return cmd.Usage()
		},
	}

	// The `options` command needs write its output to the `out` stream
	// (typically stdout). Without calling SetOutput here, the Usage()
	// function call will fall back to stderr.
	//
	// See https://github.com/kubernetes/kubernetes/pull/46394 for details.
	cmd.SetOut(ioStreams.Out())

	cmdgroups.UseOptionsTemplates(cmd)
	return cmd
}

type GlobalOptions struct {
	nonInteractive *bool
}

func NewGlobalOptions() *GlobalOptions {
	opts := &GlobalOptions{
		nonInteractive: BoolPtr(true),
	}

	return opts
}

func (g GlobalOptions) FlagSet() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{
		Usage:                nil,
		SortFlags:            false,
		ParseErrorsWhitelist: pflag.ParseErrorsWhitelist{},
	}

	flagSet.BoolVar(g.nonInteractive, "non-interactive", true, "Disable interactive mode.")

	return flagSet
}

func (g GlobalOptions) Validate() error {
	var mErr *multierror.Error

	return mErr.ErrorOrNil()
}

func BoolPtr(val bool) *bool {
	return &val
}
