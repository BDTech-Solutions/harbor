package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "harbor",
	Short: "Harbor - Development Environment Bootstrapper",
	Long: `Harbor bootstraps Laravel and WordPress development environments.
It manages dependencies via Docker and creates bootstrap scripts,
ensuring idempotency without requiring local PHP/Composer installs.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute(version string) {
	rootCmd.AddCommand(newVersionCmd(version))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "❌ "+err.Error())
		os.Exit(1)
	}
}

func newVersionCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the Harbor version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}
