package scaffold

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyIfMissing copies a single file from an embed.FS into destDir,
// preserving the filename. Does nothing if the file already exists.
func CopyIfMissing(files fs.FS, srcPath, destDir string) error {
	filename := filepath.Base(srcPath)
	dest := filepath.Join(destDir, filename)

	if _, err := os.Stat(dest); err == nil {
		return nil // already exists, idempotent
	}

	data, err := fs.ReadFile(files, srcPath)
	if err != nil {
		return fmt.Errorf("embedded file %q not found: %w", srcPath, err)
	}

	if err := os.WriteFile(dest, data, 0644); err != nil {
		return fmt.Errorf("could not write %s: %w", dest, err)
	}

	return nil
}

// CopyExecutableIfMissing is like CopyIfMissing but sets the file as executable (0755).
// Use for scripts like harbor.sh.
func CopyExecutableIfMissing(files fs.FS, srcPath, destDir string) error {
	filename := filepath.Base(srcPath)
	dest := filepath.Join(destDir, filename)

	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	data, err := fs.ReadFile(files, srcPath)
	if err != nil {
		return fmt.Errorf("embedded file %q not found: %w", srcPath, err)
	}

	if err := os.WriteFile(dest, data, 0755); err != nil {
		return fmt.Errorf("could not write %s: %w", dest, err)
	}

	return nil
}

// MakeDirs creates all given directories, including parents.
func MakeDirs(dirs ...string) error {
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("could not create directory %s: %w", d, err)
		}
	}
	return nil
}
