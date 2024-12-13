package repository

import (
	"context"

	"github.com/artarts36/depexplorer"
)

type Explorer interface {
	ExploreRepository(ctx context.Context, repo Repo) (map[depexplorer.DependencyManager]*depexplorer.File, error)
}
