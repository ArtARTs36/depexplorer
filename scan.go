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
	Read() ([]byte, error)
}

type dirProjectFileIterator struct {
	dir string

	items []os.DirEntry
	index int
	path  string
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
	for i.index < len(i.items) {
		i.index++

		item := i.items[i.index]
		if item.IsDir() {
			continue
		}

		i.path = filepath.Join(i.dir, item.Name())

		return i.path, nil
	}

	return "", io.EOF
}

func (i *dirProjectFileIterator) Read() ([]byte, error) {
	if i.index >= len(i.items) {
		return nil, io.EOF
	}

	return os.ReadFile(i.path)
}

func ScanProjectDir(dir string) (*File, error) {
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

func ScanProject(files ProjectFileIterator) (*File, error) {
	var guessFile *guessedFile

	for path, err := files.Next(); err != io.EOF; path, err = files.Next() {
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

	fileContent, err := files.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", guessFile.Path, err)
	}

	return exploreGuessedFile(guessFile, bytesContentExplorer(fileContent))
}
