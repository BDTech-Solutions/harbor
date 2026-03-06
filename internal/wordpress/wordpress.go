package wordpress

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BDTech-Solutions/harbor/internal/docker"
)

//go:embed templates
var templates embed.FS

// Init sets up a WordPress project in dir:
// copies templates, creates directory structure, starts containers,
// and fixes file ownership so the host user can edit the files.
func Init(dir string) error {
	fmt.Println("⚓ Harbor - Initializing WordPress project")

	if err := copyTemplateIfMissing(dir, "docker-compose.yml"); err != nil {
		return err
	}
	if err := copyTemplateIfMissing(dir, ".gitignore"); err != nil {
		return err
	}
	if err := copyTemplateIfMissing(dir, ".env"); err != nil {
		return err
	}

	dirs := []string{
		filepath.Join(dir, "wp", "wp-content", "plugins"),
		filepath.Join(dir, "wp", "wp-content", "themes"),
		filepath.Join(dir, "wp", "wp-content", "uploads"),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("could not create directory %s: %w", d, err)
		}
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

	port, _ := readEnvValue(filepath.Join(dir, ".env"), "WP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("✅ WordPress project initialized!")
	fmt.Printf("🌐 Open http://localhost:%s to access the site.\n", port)
	return nil
}

// Up starts existing WordPress containers.
func Up(dir string) error {
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

	port, _ := readEnvValue(filepath.Join(dir, ".env"), "WP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("✅ WordPress environment ready!")
	fmt.Printf("🌐 Open http://localhost:%s\n", port)
	return nil
}

// Down stops WordPress containers.
func Down(dir string) error {
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

// waitForWordPress polls until wp/wp-settings.php appears, then fixes ownership.
func waitForWordPress(dir string) error {
	fmt.Println("⏳ Waiting for WordPress to extract its files...")

	wpSettings := filepath.Join(dir, "wp", "wp-settings.php")
	timeout := 30 * time.Second
	deadline := time.Now().Add(timeout)

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

// copyTemplateIfMissing copies an embedded template to dir if the file doesn't exist yet.
func copyTemplateIfMissing(dir, filename string) error {
	dest := filepath.Join(dir, filename)
	if _, err := os.Stat(dest); err == nil {
		return nil // already exists
	}

	data, err := fs.ReadFile(templates, "templates/"+filename)
	if err != nil {
		return fmt.Errorf("embedded template %q not found: %w", filename, err)
	}

	if err := os.WriteFile(dest, data, 0644); err != nil {
		return fmt.Errorf("could not write %s: %w", filename, err)
	}

	return nil
}

// readEnvValue parses KEY=VALUE lines from a .env file.
func readEnvValue(envFile, key string) (string, error) {
	f, err := os.Open(envFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	prefix := key + "="
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, prefix) {
			return strings.TrimPrefix(line, prefix), nil
		}
	}
	return "", fmt.Errorf("key %q not found in %s", key, envFile)
}
