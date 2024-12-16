package e2etests

import (
	"context"
	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/github"
	repository_slog "github.com/artarts36/depexplorer/pkg/repositoryslog"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/artarts36/depexplorer/pkg/repository"
)

func TestGithubExploreRepository(t *testing.T) {
	explorer := repository.NewExplorerWithLogger(github.NewClient(nil), repository_slog.New())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deps, err := explorer.ExploreRepository(ctx, repository.Repo{
		Owner: "artarts36",
		Name:  "depexplorer",
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, deps[depexplorer.DependencyManagerGoMod])
}
