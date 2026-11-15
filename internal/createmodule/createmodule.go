package createmodule

import (
	"fmt"
	"os"
)

func Run(name string) error {
	if err := os.MkdirAll(name, 0o755); err != nil {
		return fmt.Errorf("could not create module %q: %w", name, err)
	}

	return nil
}
