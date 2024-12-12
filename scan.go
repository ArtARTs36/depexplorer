package depexplorer

import (
	"errors"
	"fmt"
	"io"
	"os"

	"path/filepath"
)

type ProjectFileIterator interface {
	Next() (string, error) // filepath, error
	Read(filepath string) ([]byte, error)
}

type dirProjectFileIterator struct {
	dir string

	items []os.DirEntry
	index int
}

func newDirProjectFileIterator(dir string) (*dirProjectFileIterator, error) {
	items, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	it := &dirProjectFileIterator{dir: dir, items: items, index: -1}
	return it, nil
}

func (i *dirProjectFileIterator) Next() (string, error) {
	for i.index < len(i.items)-1 {
		i.index++

		item := i.items[i.index]
		if item.IsDir() {
			continue
		}

		return filepath.Join(i.dir, item.Name()), nil
	}

	return "", io.EOF
}

func (i *dirProjectFileIterator) Read(filepath string) ([]byte, error) {
	return os.ReadFile(filepath)
}

func ScanProjectDir(dir string) (map[DependencyManager]*File, error) {
	it, err := newDirProjectFileIterator(dir)
	if err != nil {
		return nil, err
	}

	return ScanProject(it)
}

func bytesContentExplorer(bytes []byte) fileContentExplorer {
	return func(_ string, explorer FileExplorer) (*File, error) {
		return explorer(bytes)
	}
}

func ScanProject(files ProjectFileIterator) (map[DependencyManager]*File, error) {
	guessedFiles := map[DependencyManager]*guessedFile{}

	for path, err := files.Next(); err != io.EOF; path, err = files.Next() {
		guessFile, guessErr := guess(path)
		if guessErr != nil {
			continue
		}

		if gf, exists := guessedFiles[guessFile.DependencyManager]; !exists || !gf.IsLockFile {
			guessedFiles[guessFile.DependencyManager] = guessFile
		}
	}

	if len(guessedFiles) == 0 {
		return nil, errors.New("no dependency files found")
	}

	depFiles := map[DependencyManager]*File{}
	for _, guessFile := range guessedFiles {
		fileContent, err := files.Read(guessFile.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %q: %w", guessFile.Path, err)
		}

		depFile, err := exploreGuessedFile(guessFile, bytesContentExplorer(fileContent))
		if err != nil {
			return nil, err
		}

		depFiles[depFile.DependencyManager] = depFile
	}

	return depFiles, nil
}
