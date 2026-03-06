package cmd

import (
	"fmt"
	"os"

	"github.com/BDTech-Solutions/harbor/internal/wordpress"
	"github.com/spf13/cobra"
)

var wordpressCmd = &cobra.Command{
	Use:   "wordpress",
	Short: "Manage WordPress projects",
	Long:  `Commands to initialize and manage WordPress development environments via Docker.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("subcommand required: init | up | down")
	},
}

var wordpressInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new WordPress project in the current directory",
	Long: `Sets up Docker Compose, .env, .gitignore and directory structure
for a WordPress project. Starts containers automatically on first run.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current directory: %w", err)
		}
		return wordpress.Init(cwd)
	},
}

var wordpressUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Start WordPress containers",
	Long:  `Runs 'docker compose up -d' in the current project directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current directory: %w", err)
		}
		return wordpress.Up(cwd)
	},
}

var wordpressDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop WordPress containers",
	Long:  `Runs 'docker compose down' in the current project directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current directory: %w", err)
		}
		return wordpress.Down(cwd)
	},
}

func init() {
	wordpressCmd.AddCommand(wordpressInitCmd)
	wordpressCmd.AddCommand(wordpressUpCmd)
	wordpressCmd.AddCommand(wordpressDownCmd)
	rootCmd.AddCommand(wordpressCmd)
}
