package github

import (
	"context"
	"fmt"
	"io"

	"github.com/google/go-github/github"
)

type fileIterator struct {
	repository Repository
	ctx        context.Context
	files      []*github.RepositoryContent
	client     *github.Client
	logger     Logger

	index int
}

func newFileIterator(
	repository Repository,
	ctx context.Context,
	files []*github.RepositoryContent,
	client *github.Client,
	logger Logger,
) *fileIterator {
	return &fileIterator{
		repository: repository,
		ctx:        ctx,
		files:      files,
		client:     client,
		index:      -1,
		logger:     logger,
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

func (i *fileIterator) Read() ([]byte, error) {
	if i.index >= len(i.files)-1 {
		return nil, io.EOF
	}

	filepath := i.files[i.index].Path
	if filepath == nil {
		return nil, fmt.Errorf("file with index %d must have filepath", i.index)
	}

	i.logger("get file contents", map[string]interface{}{
		"repo_owner":    i.repository.Owner,
		"repo_name":     i.repository.Repo,
		"repo_filepath": *filepath,
	})

	file, _, _, err := i.client.Repositories.GetContents(i.ctx, i.repository.Owner, i.repository.Repo, *filepath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get contents for file %q: %v", *filepath, err)
	}

	if file.Content == nil {
		return nil, fmt.Errorf("file with index %d must have content", i.index)
	}

	return []byte(*file.Content), nil
}
