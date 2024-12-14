package repository

import (
	"context"

	"github.com/artarts36/depexplorer"
)

type Client interface {
	ListFiles(ctx context.Context, repo Repo, dir string) (depexplorer.DirectoryFileIterator, error)
	//ListUserRepositories(ctx context.Context, user string) ([]Repo, error)
	//ListOrgRepositories(ctx context.Context, org string) ([]Repo, error)
}
