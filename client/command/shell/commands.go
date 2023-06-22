package shell

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/bishopfox/sliver/client/command/flags"
	"github.com/bishopfox/sliver/client/command/help"
	"github.com/bishopfox/sliver/client/console"
	consts "github.com/bishopfox/sliver/client/constants"
)

// Commands returns the “ command and its subcommands.
func Commands(con *console.SliverConsoleClient) []*cobra.Command {
	shellCmd := &cobra.Command{
		Use:   consts.ShellStr,
		Short: "Start an interactive shell",
		Long:  help.GetHelpFor([]string{consts.ShellStr}),
		Run: func(cmd *cobra.Command, args []string) {
			ShellCmd(cmd, con, args)
		},
		GroupID:     consts.ExecutionHelpGroup,
		Annotations: flags.RestrictTargets(consts.SessionCmdsFilter),
	}
	flags.Bind("", false, shellCmd, func(f *pflag.FlagSet) {
		f.BoolP("no-pty", "y", false, "disable use of pty on macos/linux")
		f.StringP("shell-path", "s", "", "path to shell interpreter")

		f.Int64P("timeout", "t", flags.DefaultTimeout, "grpc timeout in seconds")
	})

	return []*cobra.Command{shellCmd}
}
