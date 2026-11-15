package createproject

import (
	"fmt"
	"os"
)

func Run(name string) error {
	if _, err := os.Stat(name); err == nil {
		return fmt.Errorf("project %q already exists", name)
	}

	if err := os.MkdirAll(name, 0o755); err != nil {
		return fmt.Errorf("could not create project %q: %w", name, err)
	}

	return nil
}
