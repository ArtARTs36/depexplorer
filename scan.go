package depexplorer

import (
	"errors"
	"os"
	"path/filepath"
)

type File struct {
	Path string
	Name string

	DependencyManager DependencyManager
	Dependencies      []*Dependency
}

func ScanProject(dir string) (*File, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var guessFile *guessedFile

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(dir, entry.Name())

		var guessErr error
		guessFile, guessErr = guess(path)
		if guessErr != nil {
			continue
		}

		if guessFile.IsLockFile || !guessFile.CanHaveLockFile {
			break
		}
	}

	if guessFile == nil {
		return nil, errors.New("no dependency files found")
	}

	return exploreGuessedFile(guessFile)
}
