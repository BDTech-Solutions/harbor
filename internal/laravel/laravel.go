package laravel

import (
	"embed"
	"fmt"
	"os"

	"github.com/BDTech-Solutions/harbor/internal/docker"
	"github.com/BDTech-Solutions/harbor/internal/scaffold"
)

//go:embed templates/harbor.sh
var templates embed.FS

// Stack implements stack.Stack for Laravel projects.
type Stack struct{}

func (s Stack) Init(dir string) error {
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

	if err := scaffold.CopyExecutableIfMissing(templates, "templates/harbor.sh", dir); err != nil {
		return err
	}

	fmt.Println("✅ Laravel project created with Sail.")
	fmt.Println("👉 Run ./harbor.sh to start the environment.")
	return nil
}

func (s Stack) Up(dir string) error {
	fmt.Println("🚀 Starting Laravel environment with Sail...")

	if err := assertLaravelProject(dir); err != nil {
		return err
	}

	if err := docker.Check(); err != nil {
		return err
	}

	return docker.Run(dir, "vendor/bin/sail", "up", "-d")
}

func (s Stack) Down(dir string) error {
	fmt.Println("🛑 Stopping Laravel environment...")

	if err := assertLaravelProject(dir); err != nil {
		return err
	}

	if err := docker.Check(); err != nil {
		return err
	}

	return docker.Run(dir, "vendor/bin/sail", "down")
}

// Bootstrap copies harbor.sh into an existing Laravel project.
// This is Laravel-specific behaviour with no equivalent in other stacks,
// so it lives outside the Stack interface.
func Bootstrap(dir string) error {
	fmt.Println("🔎 Checking Laravel project...")

	if err := assertLaravelProject(dir); err != nil {
		return err
	}

	if err := scaffold.CopyExecutableIfMissing(templates, "templates/harbor.sh", dir); err != nil {
		return err
	}

	fmt.Println("✅ harbor.sh created successfully.")
	fmt.Println("👉 Run ./harbor.sh to start the Laravel environment.")
	return nil
}

func assertLaravelProject(dir string) error {
	required := []string{"artisan", "composer.json", "bootstrap/app.php"}
	for _, f := range required {
		path := dir + "/" + f
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("this directory does not look like a Laravel project (missing %s)", f)
		}
	}
	return nil
}
