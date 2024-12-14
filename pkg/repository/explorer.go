package repository

import (
	"context"
	"fmt"

	"github.com/artarts36/depexplorer"
)

type Explorer struct {
	client Client
	logger Logger
}

func NewExplorer(client Client, logger Logger) *Explorer {
	return &Explorer{client: client, logger: logger}
}

func (e *Explorer) ExploreRepository(
	ctx context.Context,
	repo Repo,
	opts *ExploreOpts,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	e.logger("listing repo files", map[string]interface{}{
		"repo_owner": repo.Owner,
		"repo_name":  repo.Name,
	})

	iterator, err := e.client.ListFiles(ctx, repo, opts.getDirectory())
	if err != nil {
		return nil, fmt.Errorf("failed to list repo files: %w", err)
	}

	return depexplorer.ExploreDirectory(newLogIterator(e.logger, iterator, repo))
}
