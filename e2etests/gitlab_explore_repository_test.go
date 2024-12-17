package e2etests

import (
	"context"
	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/gitlab"
	repository_slog "github.com/artarts36/depexplorer/pkg/repositoryslog"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/artarts36/depexplorer/pkg/repository"
)

func TestGitlabExploreRepository(t *testing.T) {
	gitlabClient, err := gitlab.NewClientWithToken("", nil)

	explorer := repository.NewExplorerWithLogger(gitlabClient, repository_slog.New())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deps, err := explorer.ExploreRepository(ctx, repository.Repo{
		Owner: "gitlab-org/api",
		Name:  "client-go",
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, deps[depexplorer.DependencyManagerGoMod])
}
