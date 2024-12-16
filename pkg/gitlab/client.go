package gitlab

import (
	"context"
	"fmt"
	"github.com/artarts36/depexplorer"
	"github.com/artarts36/depexplorer/pkg/repository"
	"gitlab.com/gitlab-org/api/client-go"
)

type Client struct {
	client *gitlab.Client
}

func NewClient(client *gitlab.Client) *Client {
	return &Client{
		client: client,
	}
}

func NewClientWithToken(token string) (*Client, error) {
	client, err := gitlab.NewClient(token)
	if err != nil {
		return nil, err
	}
	return NewClient(client), nil
}

func (c *Client) ListFiles(
	ctx context.Context,
	repo repository.Repo,
	opts repository.ListRepoFilesOpts,
) (depexplorer.DirectoryFileIterator, error) {
	projectID := c.createProjectID(repo)

	listTreeOpts := &gitlab.ListTreeOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
		Path: &opts.Directory,
	}
	ref := opts.Ref
	if ref == "" {
		var err error
		ref, err = c.getDefaultBranch(ctx, projectID)
		if err != nil {
			return nil, fmt.Errorf("failed to default branch: %w", err)
		}
	}
	listTreeOpts.Ref = &ref

	nodes, _, err := c.client.Repositories.ListTree(projectID, listTreeOpts, gitlab.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("list project tree: %w", err)
	}

	return newFileIterator(nodes, func(path string) (*gitlab.File, error) {
		node, _, fErr := c.client.RepositoryFiles.GetFile(projectID, "go.mod", &gitlab.GetFileOptions{
			Ref: &ref,
		}, gitlab.WithContext(ctx))

		return node, fErr
	}), nil
}

func (c *Client) getDefaultBranch(ctx context.Context, projectID string) (string, error) {
	branches, _, err := c.client.Branches.ListBranches(projectID, &gitlab.ListBranchesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
	}, gitlab.WithContext(ctx))
	if err != nil {
		return "", fmt.Errorf("failed to list repository branches: %w", err)
	}

	defaultBranch := ""
	for _, branch := range branches {
		if branch.Default {
			defaultBranch = branch.Name
		}
	}

	if defaultBranch == "" {
		return "", fmt.Errorf("default branch not found")
	}

	return defaultBranch, nil
}

func (c *Client) createProjectID(repo repository.Repo) string {
	return fmt.Sprintf("%s/%s", repo.Owner, repo.Name)
}
