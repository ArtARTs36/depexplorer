package github_test

import (
	"context"

	"testing"

	"github.com/stretchr/testify/require"

	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/github"
	"github.com/artarts36/depexplorer/pkg/repository"
)

func TestScanThisRepository_Integration(t *testing.T) {
	files, err := github.ExploreRepository(
		context.Background(),
		repository.Repo{
			Owner: "artarts36",
			Name:  "depexplorer",
		},
	)
	require.NoError(t, err)
	require.NotNil(t, files[depexplorer.DependencyManagerGoMod])
}
