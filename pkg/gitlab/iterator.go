package gitlab

import (
	"encoding/base64"
	"fmt"
	"io"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type fileLoader func(path string) (*gitlab.File, error)

type fileIterator struct {
	files      []*gitlab.TreeNode
	fileLoader fileLoader

	index int
}

func newFileIterator(
	files []*gitlab.TreeNode,
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

	return i.files[i.index].Path, nil
}

func (i *fileIterator) Read(filepath string) ([]byte, error) {
	file, err := i.fileLoader(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load file: %w", err)
	}

	fileContent, err := getFileContent(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode file %q: %v", filepath, err)
	}

	return fileContent, nil
}

func getFileContent(file *gitlab.File) ([]byte, error) {
	switch file.Encoding {
	case "base64":
		c, err := base64.StdEncoding.DecodeString(file.Content)
		return c, err
	case "":
		return []byte(file.Content), nil
	default:
		return nil, fmt.Errorf("unsupported content encoding: %v", file.Encoding)
	}
}
