package repository_test

import (
	"github.com/artarts36/depexplorer/pkg/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepoFromURI(t *testing.T) {
	cases := []struct {
		Title    string
		URI      string
		Expected repository.Repo
	}{
		{
			Title: "parse simple github repository",
			URI:   "https://github.com/ArtARTs36/depexplorer",
			Expected: repository.Repo{
				Host:  "github.com",
				Owner: "ArtARTs36",
				Name:  "depexplorer",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			repo, err := repository.NewRepoFromURI(c.URI)
			require.NoError(t, err)
			assert.Equal(t, c.Expected, repo)
		})
	}
}
