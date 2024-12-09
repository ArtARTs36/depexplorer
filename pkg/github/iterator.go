package github

import (
	"fmt"
	"io"

	"github.com/google/go-github/v67/github"
)

type fileLoader func(path string) (*github.RepositoryContent, error)

type fileIterator struct {
	files      []*github.RepositoryContent
	fileLoader fileLoader

	index int
}

func newFileIterator(
	files []*github.RepositoryContent,
	fileLoader fileLoader,
) *fileIterator {
	return &fileIterator{
		files:      files,
		index:      -1,
		fileLoader: fileLoader,
	}
}

func (i *fileIterator) Next() (string, error) {
	i.index++
	if i.index >= len(i.files)-1 {
		return "", io.EOF
	}

	filename := i.files[i.index].Name
	if filename == nil {
		return "", fmt.Errorf("file with index %d must have filename", i.index)
	}

	return *filename, nil
}

func (i *fileIterator) Read(filepath string) ([]byte, error) {
	file, err := i.fileLoader(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load file: %w", err)
	}

	fileContent, err := file.GetContent()
	if err != nil {
		return nil, fmt.Errorf("failed to decode file %q: %v", filepath, err)
	}

	return []byte(fileContent), nil
}
