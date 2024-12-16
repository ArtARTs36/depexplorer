package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v67/github"

	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/repository"
)

type Client struct {
	client *github.Client
}

func NewClient(client *github.Client) *Client {
	if client == nil {
		client = github.NewClient(nil)
	}

	return &Client{client: client}
}

func NewClientWithToken(token string) *Client {
	return NewClient(github.NewClient(nil).WithAuthToken(token))
}

func (c *Client) ListFiles(
	ctx context.Context,
	repo repository.Repo,
	opts repository.ListRepoFilesOpts,
) (depexplorer.DirectoryFileIterator, error) {
	githubOpts := &github.RepositoryContentGetOptions{
		Ref: opts.Ref,
	}

	_, files, _, err := c.client.Repositories.GetContents(
		ctx,
		repo.Owner,
		repo.Name,
		opts.Directory,
		githubOpts,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get contents from repo: %w", err)
	}

	return newFileIterator(files, func(path string) (*github.RepositoryContent, error) {
		return c.getFile(ctx, repo, path, opts)
	}), nil
}

func (c *Client) getFile(
	ctx context.Context,
	repo repository.Repo,
	path string,
	opts repository.ListRepoFilesOpts,
) (*github.RepositoryContent, error) {
	githubOpts := &github.RepositoryContentGetOptions{
		Ref: opts.Ref,
	}

	file, _, _, err := c.client.Repositories.GetContents(ctx, repo.Owner, repo.Name, path, githubOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get contents for file %q: %v", path, err)
	}

	return file, nil
}
