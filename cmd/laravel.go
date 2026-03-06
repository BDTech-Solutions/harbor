package cmd

import (
	"fmt"
	"os"

	"github.com/BDTech-Solutions/harbor/internal/laravel"
	"github.com/spf13/cobra"
)

// NewLaravelCommand builds the "harbor laravel" command tree.
func NewLaravelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "laravel",
		Short: "Manage Laravel projects",
		Long:  `Commands to initialize and bootstrap Laravel development environments.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(newLaravelInitCommand())
	cmd.AddCommand(newLaravelBootstrapCommand())

	return cmd
}

func newLaravelInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create a new Laravel project in the current directory",
		Long: `Creates a new Laravel project using Docker + Composer.
Installs Laravel Sail and generates a harbor.sh bootstrap script.
The current directory must be empty.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get current directory: %w", err)
			}
			return laravel.Init(cwd)
		},
	}
}

func newLaravelBootstrapCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "bootstrap",
		Short: "Generate harbor.sh for an existing Laravel project",
		Long: `Copies the harbor.sh bootstrap template into the current Laravel project.
Useful for cloned projects that don't have the script yet.
The current directory must contain a valid Laravel project.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get current directory: %w", err)
			}
			return laravel.Bootstrap(cwd)
		},
	}
}
