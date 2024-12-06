package depexplorer

import (
	"fmt"
	"os"
)

type FileExplorer func(file []byte) ([]*Dependency, error)

func Explore(path string, explorer FileExplorer) ([]*Dependency, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return explorer(file)
}
