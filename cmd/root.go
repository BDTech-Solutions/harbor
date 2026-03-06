package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewRootCommand builds and returns the root harbor command.
// All subcommands are registered here explicitly — no init() magic.
func NewRootCommand(version string) *cobra.Command {
	root := &cobra.Command{
		Use:   "harbor",
		Short: "Harbor - Development Environment Bootstrapper",
		Long: `Harbor bootstraps Laravel and WordPress development environments.
It manages dependencies via Docker and creates bootstrap scripts,
ensuring idempotency without requiring local PHP/Composer installs.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	root.AddCommand(newVersionCommand(version))
	root.AddCommand(NewLaravelCommand())
	root.AddCommand(NewWordpressCommand())

	return root
}

func newVersionCommand(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the Harbor version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, version)
		},
	}
}
