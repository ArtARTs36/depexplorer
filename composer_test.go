package depexplorer_test

import (
	"github.com/artarts36/depexplorer"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExploreComposer(t *testing.T) {
	file := `{
    "name": "artarts36/local-file-system",
    "require": {
        "artarts36/file-system-contracts": "^0.2.0"
    },
    "require-dev": {
        "phpunit/phpunit": "^9.5"
    }
}`

	expected := []*depexplorer.Dependency{
		{
			Name:    "artarts36/file-system-contracts",
			Version: depexplorer.Version{Full: "^0.2.0"},
		},
		{
			Name:    "phpunit/phpunit",
			Version: depexplorer.Version{Full: "^9.5"},
		},
	}

	got, err := depexplorer.ExploreComposerJSON([]byte(file))
	require.NoError(t, err)
	require.Equal(t, expected, got.Dependencies)
}
