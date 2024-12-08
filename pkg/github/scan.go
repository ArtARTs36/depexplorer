package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"

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
	logger.Printf("listing repository files", map[string]string{
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

	return depexplorer.ScanProject(newFileIterator(repository, ctx, files, client, logger))
}
