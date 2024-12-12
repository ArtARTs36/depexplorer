package github_test

import (
	"context"

	"testing"

	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/github"
	githubClient "github.com/google/go-github/v67/github"
	"github.com/stretchr/testify/require"
)

func TestScanThisRepository_Integration(t *testing.T) {
	client := githubClient.NewClient(nil)

	files, err := github.ScanRepository(
		context.Background(),
		client,
		github.Repository{
			Owner: "artarts36",
			Repo:  "depexplorer",
		},
		github.NoopLogger(),
	)
	require.NoError(t, err)
	require.NotNil(t, files[depexplorer.DependencyManagerGoMod])
}
