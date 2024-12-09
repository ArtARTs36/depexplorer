package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v67/github"

	"github.com/artarts36/depexplorer"
)

type Repository struct {
	Owner string
	Repo  string
}

func ScanRepository(
	ctx context.Context,
	client *github.Client,
	repository Repository,
	logger Logger,
) (*depexplorer.File, error) {
	logger("listing repository files", map[string]interface{}{
		"repo_owner": repository.Owner,
		"repo_name":  repository.Repo,
	})

	_, files, _, err := client.Repositories.GetContents(
		ctx,
		repository.Owner,
		repository.Repo,
		"",
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list repository files: %w", err)
	}

	loadFile := func(filepath string) (*github.RepositoryContent, error) {
		logger("get file contents", map[string]interface{}{
			"repo_owner":    repository.Owner,
			"repo_name":     repository.Repo,
			"repo_filepath": filepath,
		})

		file, _, _, fErr := client.Repositories.GetContents(ctx, repository.Owner, repository.Repo, filepath, nil)
		if fErr != nil {
			return nil, fmt.Errorf("failed to get contents for file %q: %v", filepath, err)
		}

		return file, nil
	}

	return depexplorer.ScanProject(newFileIterator(files, loadFile))
}
