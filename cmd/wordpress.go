package cmd

import (
	"github.com/BDTech-Solutions/harbor/internal/wordpress"
	"github.com/spf13/cobra"
)

func NewWordpressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wordpress",
		Short: "Manage WordPress projects",
		Long:  `Commands to initialize and manage WordPress development environments via Docker.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	stack := wordpress.Stack{}

	cmd.AddCommand(newWordpressInitCommand(stack))
	cmd.AddCommand(newWordpressUpCommand(stack))
	cmd.AddCommand(newWordpressDownCommand(stack))

	return cmd
}

func newWordpressInitCommand(stack wordpress.Stack) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a new WordPress project in the current directory",
		Long: `Sets up Docker Compose, .env, .gitignore and directory structure
for a WordPress project. Starts containers automatically on first run.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return stack.Init(cwd())
		},
	}
}

func newWordpressUpCommand(stack wordpress.Stack) *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Start WordPress containers",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return stack.Up(cwd())
		},
	}
}

func newWordpressDownCommand(stack wordpress.Stack) *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Stop WordPress containers",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return stack.Down(cwd())
		},
	}
}
