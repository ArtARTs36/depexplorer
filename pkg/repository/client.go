package repository

import (
	"context"

	"github.com/artarts36/depexplorer"
)

type Client interface {
	ListFiles(ctx context.Context, repo Repo, opts ListRepoFilesOpts) (depexplorer.DirectoryFileIterator, error)
	// ListUserRepositories(ctx context.Context, user string) ([]Repo, error)
	// ListOrgRepositories(ctx context.Context, org string) ([]Repo, error)
}

type ListRepoFilesOpts struct {
	Directory string
	Ref       string
}
