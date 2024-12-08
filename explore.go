package depexplorer

import (
	"fmt"
	"os"
)

type File struct {
	Path string
	Name string

	DependencyManager DependencyManager
	Dependencies      []*Dependency
}

type FileExplorer func(file []byte) (*File, error)

func Explore(path string, explorer FileExplorer) (*File, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return explorer(file)
}
