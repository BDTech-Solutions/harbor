package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read parses a .env file and returns all key-value pairs.
func Read(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %w", path, err)
	}
	defer f.Close()

	values := make(map[string]string)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip blank lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		values[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}

	return values, scanner.Err()
}

// Get returns a single value from a .env file, with a fallback if missing.
func Get(path, key, fallback string) string {
	values, err := Read(path)
	if err != nil {
		return fallback
	}
	if v, ok := values[key]; ok && v != "" {
		return v
	}
	return fallback
}
