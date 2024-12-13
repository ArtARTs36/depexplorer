package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v67/github"

	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/repository"
)

type Explorer struct {
	client *github.Client
	logger Logger
}

func NewExplorer(client *github.Client, logger Logger) *Explorer {
	return &Explorer{client: client, logger: logger}
}

func (e *Explorer) ExploreRepository(
	ctx context.Context,
	repository repository.Repo,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	e.logger("listing repository files", map[string]interface{}{
		"repo_owner": repository.Owner,
		"repo_name":  repository.Name,
	})

	_, files, _, err := e.client.Repositories.GetContents(
		ctx,
		repository.Owner,
		repository.Name,
		repository.Directory,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list repository files: %w", err)
	}

	loadFile := func(filepath string) (*github.RepositoryContent, error) {
		e.logger("get file contents", map[string]interface{}{
			"repo_owner":    repository.Owner,
			"repo_name":     repository.Name,
			"repo_filepath": filepath,
		})

		file, _, _, fErr := e.client.Repositories.GetContents(ctx, repository.Owner, repository.Name, filepath, nil)
		if fErr != nil {
			return nil, fmt.Errorf("failed to get contents for file %q: %v", filepath, err)
		}

		return file, nil
	}

	return depexplorer.ExploreDirectory(newFileIterator(files, loadFile))
}
