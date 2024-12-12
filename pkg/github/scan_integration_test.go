package github_test

import (
	"context"

	"testing"

	"github.com/stretchr/testify/require"

	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/github"
)

func TestScanThisRepository_Integration(t *testing.T) {
	files, err := github.ScanRepository(
		context.Background(),
		github.Repository{
			Owner: "artarts36",
			Repo:  "depexplorer",
		},
		nil,
	)
	require.NoError(t, err)
	require.NotNil(t, files[depexplorer.DependencyManagerGoMod])
}
