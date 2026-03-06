package wordpress

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BDTech-Solutions/harbor/internal/docker"
	"github.com/BDTech-Solutions/harbor/internal/env"
	"github.com/BDTech-Solutions/harbor/internal/scaffold"
)

//go:embed templates
var templates embed.FS

const defaultPort = "8080"

// Stack implements stack.Stack for WordPress projects.
type Stack struct{}

func (s Stack) Init(dir string) error {
	fmt.Println("⚓ Harbor - Initializing WordPress project")

	files := []string{"docker-compose.yml", ".gitignore", ".env"}
	for _, f := range files {
		if err := scaffold.CopyIfMissing(templates, "templates/"+f, dir); err != nil {
			return err
		}
	}

	err := scaffold.MakeDirs(
		filepath.Join(dir, "wp", "wp-content", "plugins"),
		filepath.Join(dir, "wp", "wp-content", "themes"),
		filepath.Join(dir, "wp", "wp-content", "uploads"),
	)
	if err != nil {
		return err
	}

	fmt.Println("✅ File structure created.")

	if err := docker.Check(); err != nil {
		return err
	}

	fmt.Println("🐳 Starting containers...")
	if err := docker.Compose(dir, "up", "-d"); err != nil {
		return fmt.Errorf("docker compose up failed: %w", err)
	}

	if err := waitForWordPress(dir); err != nil {
		fmt.Println("⚠️  Warning:", err.Error())
	}

	port := env.Get(filepath.Join(dir, ".env"), "WP_PORT", defaultPort)
	fmt.Println("✅ WordPress project initialized!")
	fmt.Printf("🌐 Open http://localhost:%s to access the site.\n", port)
	return nil
}

func (s Stack) Up(dir string) error {
	fmt.Println("🚀 Starting WordPress containers...")

	if err := docker.Check(); err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(dir, "docker-compose.yml")); os.IsNotExist(err) {
		return fmt.Errorf("docker-compose.yml not found — run 'harbor wordpress init' first")
	}

	if err := docker.Compose(dir, "up", "-d"); err != nil {
		return fmt.Errorf("docker compose up failed: %w", err)
	}

	port := env.Get(filepath.Join(dir, ".env"), "WP_PORT", defaultPort)
	fmt.Println("✅ WordPress environment ready!")
	fmt.Printf("🌐 Open http://localhost:%s\n", port)
	return nil
}

func (s Stack) Down(dir string) error {
	fmt.Println("🛑 Stopping WordPress containers...")

	if err := docker.Check(); err != nil {
		return err
	}

	if err := docker.Compose(dir, "down"); err != nil {
		return fmt.Errorf("docker compose down failed: %w", err)
	}

	fmt.Println("✅ WordPress containers stopped.")
	return nil
}

func waitForWordPress(dir string) error {
	fmt.Println("⏳ Waiting for WordPress to extract its files...")

	wpSettings := filepath.Join(dir, "wp", "wp-settings.php")
	deadline := time.Now().Add(30 * time.Second)

	for time.Now().Before(deadline) {
		if _, err := os.Stat(wpSettings); err == nil {
			fmt.Printf("🐳 Fixing file ownership for user %d...\n", os.Getuid())
			uid := fmt.Sprintf("%d:%d", os.Getuid(), os.Getgid())
			_ = docker.Run(dir,
				"compose", "exec", "-u", "root", "wordpress",
				"chown", "-R", uid, "/var/www/html",
			)
			fmt.Println("✅ Permissions adjusted.")
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("WordPress took too long to initialize — check 'docker compose logs'")
}
