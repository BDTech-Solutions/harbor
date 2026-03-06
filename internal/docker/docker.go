package docker

import (
	"fmt"
	"os"
	"os/exec"
)

// Check verifies that Docker is installed and the daemon is running.
func Check() error {
	if _, err := exec.LookPath("docker"); err != nil {
		return fmt.Errorf("docker not found — please install Docker before continuing")
	}

	cmd := exec.Command("docker", "info")
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker daemon is not running — start Docker and try again")
	}

	return nil
}

// Run executes a docker command in the given working directory,
// streaming stdout/stderr directly to the terminal.
func Run(cwd string, args ...string) error {
	cmd := exec.Command("docker", args...)
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// Compose executes a docker compose command in the given working directory.
func Compose(cwd string, args ...string) error {
	fullArgs := append([]string{"compose"}, args...)
	return Run(cwd, fullArgs...)
}
