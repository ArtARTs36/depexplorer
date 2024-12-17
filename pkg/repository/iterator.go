package repository

import (
	"github.com/artarts36/depexplorer"
)

type logIterator struct {
	logger Logger
	inner  depexplorer.DirectoryFileIterator
	repo   Repo
}

func newLogIterator(logger Logger, inner depexplorer.DirectoryFileIterator, repo Repo) *logIterator {
	return &logIterator{
		logger: logger,
		inner:  inner,
		repo:   repo,
	}
}

func (i *logIterator) Next() (string, error) {
	return i.inner.Next()
}

func (i *logIterator) Read(path string) ([]byte, error) {
	i.logger("get file contents", map[string]interface{}{
		"repo_owner":    i.repo.Owner,
		"repo_name":     i.repo.Name,
		"repo_filepath": path,
	})

	return i.inner.Read(path)
}
