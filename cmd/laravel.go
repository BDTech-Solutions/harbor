package cmd

import (
	"fmt"
	"os"

	"github.com/BDTech-Solutions/harbor/internal/laravel"
	"github.com/spf13/cobra"
)

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

	stack := laravel.Stack{}

	cmd.AddCommand(newLaravelInitCommand(stack))
	cmd.AddCommand(newLaravelUpCommand(stack))
	cmd.AddCommand(newLaravelDownCommand(stack))
	cmd.AddCommand(newLaravelBootstrapCommand())

	return cmd
}

func newLaravelInitCommand(stack laravel.Stack) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create a new Laravel project in the current directory",
		Long: `Creates a new Laravel project using Docker + Composer.
Installs Laravel Sail and generates a harbor.sh bootstrap script.
The current directory must be empty.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return stack.Init(cwd())
		},
	}
}

func newLaravelUpCommand(stack laravel.Stack) *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Start the Laravel environment with Sail",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return stack.Up(cwd())
		},
	}
}

func newLaravelDownCommand(stack laravel.Stack) *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Stop the Laravel environment",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return stack.Down(cwd())
		},
	}
}

func newLaravelBootstrapCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "bootstrap",
		Short: "Generate harbor.sh for an existing Laravel project",
		Long: `Copies the harbor.sh bootstrap template into the current Laravel project.
Useful for cloned projects that don't have the script yet.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return laravel.Bootstrap(cwd())
		},
	}
}

// cwd returns the current working directory or exits on failure.
// Defined here to avoid repeating os.Getwd() error handling in every command.
func cwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "❌ could not get current directory:", err)
		os.Exit(1)
	}
	return dir
}
