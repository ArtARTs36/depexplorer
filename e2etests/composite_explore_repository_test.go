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

func TestCompositeExploreRepository(t *testing.T) {
	explorer := repository.NewExplorerWithLogger(repository.NewClientComposite(map[string]repository.Client{
		"github.com": github.NewClient(nil),
	}), repository_slog.New())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deps, err := explorer.ExploreRepository(ctx, repository.Repo{
		Host:  "github.com",
		Owner: "artarts36",
		Name:  "depexplorer",
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, deps[depexplorer.DependencyManagerGoMod])
}
