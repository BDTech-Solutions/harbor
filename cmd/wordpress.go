package cmd

import (
	"fmt"
	"os"

	"github.com/BDTech-Solutions/harbor/internal/wordpress"
	"github.com/spf13/cobra"
)

// NewWordpressCommand builds the "harbor wordpress" command tree.
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

	cmd.AddCommand(newWordpressInitCommand())
	cmd.AddCommand(newWordpressUpCommand())
	cmd.AddCommand(newWordpressDownCommand())

	return cmd
}

func newWordpressInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a new WordPress project in the current directory",
		Long: `Sets up Docker Compose, .env, .gitignore and directory structure
for a WordPress project. Starts containers automatically on first run.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get current directory: %w", err)
			}
			return wordpress.Init(cwd)
		},
	}
}

func newWordpressUpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Start WordPress containers",
		Long:  `Runs 'docker compose up -d' in the current project directory.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get current directory: %w", err)
			}
			return wordpress.Up(cwd)
		},
	}
}

func newWordpressDownCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Stop WordPress containers",
		Long:  `Runs 'docker compose down' in the current project directory.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get current directory: %w", err)
			}
			return wordpress.Down(cwd)
		},
	}
}
