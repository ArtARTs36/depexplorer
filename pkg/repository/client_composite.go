package repository

import (
	"context"
	"fmt"

	"github.com/artarts36/depexplorer"
)

type ClientComposite struct {
	clients map[string]Client
}

func NewClientComposite(clients map[string]Client) *ClientComposite {
	return &ClientComposite{
		clients: clients,
	}
}

func (c *ClientComposite) ListFiles(
	ctx context.Context,
	repo Repo,
	opts ListRepoFilesOpts,
) (depexplorer.DirectoryFileIterator, error) {
	client, ok := c.clients[repo.Host]
	if !ok {
		return nil, fmt.Errorf("no client found for host %q", repo.Host)
	}

	return client.ListFiles(ctx, repo, opts)
}
