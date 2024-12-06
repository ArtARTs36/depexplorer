package depexplorer_test

import (
	"github.com/artarts36/depexplorer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExploreGoMod(t *testing.T) {
	gomod := `module github.com/artarts36/depexplorer

go 1.23.3

require golang.org/x/mod v0.22.0`

	expected := []*depexplorer.Dependency{
		{
			Name: "golang.org/x/mod",
			Version: depexplorer.Version{
				Full: "v0.22.0",
			},
		},
	}

	got, err := depexplorer.ExploreGoMod([]byte(gomod))
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
