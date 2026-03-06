package laravel

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/BDTech-Solutions/harbor/internal/docker"
)

//go:embed templates/harbor.sh
var templates embed.FS

// Init creates a brand-new Laravel project in dir using Docker + Composer,
// installs Sail, then drops a harbor.sh bootstrap script.
// The directory must be empty.
func Init(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("cannot read directory: %w", err)
	}
	if len(entries) > 0 {
		return fmt.Errorf("directory is not empty — aborting to avoid overwriting files")
	}

	fmt.Println("⚓ Creating Laravel project via Docker (Composer)...")

	uid := fmt.Sprintf("%d:%d", os.Getuid(), os.Getgid())

	err = docker.Run(dir,
		"run", "--rm",
		"-u", uid,
		"-v", dir+":/app",
		"-w", "/app",
		"composer:2",
		"composer", "create-project", "laravel/laravel", ".",
	)
	if err != nil {
		return fmt.Errorf("composer create-project failed: %w", err)
	}

	fmt.Println("🚢 Installing Laravel Sail...")

	err = docker.Run(dir,
		"run", "--rm",
		"-u", uid,
		"-v", dir+":/app",
		"-w", "/app",
		"laravelsail/php83-composer:latest",
		"php", "artisan", "sail:install",
	)
	if err != nil {
		return fmt.Errorf("sail:install failed: %w", err)
	}

	if err := writeHarborScript(dir); err != nil {
		return err
	}

	fmt.Println("✅ Laravel project created with Sail.")
	fmt.Println("👉 Run ./harbor.sh to start the environment.")
	return nil
}

// Bootstrap copies harbor.sh into an existing Laravel project.
// The directory must look like a Laravel project and must not already have harbor.sh.
func Bootstrap(dir string) error {
	fmt.Println("🔎 Checking Laravel project...")

	if err := assertLaravelProject(dir); err != nil {
		return err
	}

	harborScript := filepath.Join(dir, "harbor.sh")
	if _, err := os.Stat(harborScript); err == nil {
		fmt.Println("ℹ️  harbor.sh already exists in this project — nothing to do.")
		return nil
	}

	if err := writeHarborScript(dir); err != nil {
		return err
	}

	fmt.Println("✅ harbor.sh created successfully.")
	fmt.Println("👉 Run ./harbor.sh to start the Laravel environment.")
	return nil
}

// assertLaravelProject returns an error if dir doesn't look like a Laravel project.
func assertLaravelProject(dir string) error {
	required := []string{"artisan", "composer.json", filepath.Join("bootstrap", "app.php")}
	for _, f := range required {
		if _, err := os.Stat(filepath.Join(dir, f)); os.IsNotExist(err) {
			return fmt.Errorf("this directory does not look like a Laravel project (missing %s)", f)
		}
	}
	return nil
}

// writeHarborScript embeds the template and writes it to dir/harbor.sh.
func writeHarborScript(dir string) error {
	data, err := fs.ReadFile(templates, "templates/harbor.sh")
	if err != nil {
		return fmt.Errorf("embedded template not found: %w", err)
	}

	dest := filepath.Join(dir, "harbor.sh")
	if err := os.WriteFile(dest, data, 0755); err != nil {
		return fmt.Errorf("could not write harbor.sh: %w", err)
	}

	return nil
}
