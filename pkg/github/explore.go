package github

import (
	"context"

	"github.com/google/go-github/v67/github"

	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/repository"
)

var DefaultExplorer = NewExplorer(github.NewClient(nil), NoopLogger())

func ExploreRepository(
	ctx context.Context,
	repository repository.Repo,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	return DefaultExplorer.ExploreRepository(ctx, repository)
}
