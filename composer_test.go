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
		"php": "8.1"
    },
    "require-dev": {
        "phpunit/phpunit": "^9.5"
    }
}`

	expected := &depexplorer.File{
		Name:              "composer.json",
		Path:              "composer.json",
		DependencyManager: depexplorer.DependencyManagerComposer,
		Dependencies: []*depexplorer.Dependency{
			{
				Name:    "php",
				Version: depexplorer.Version{Full: "8.1"},
			},
			{
				Name:    "phpunit/phpunit",
				Version: depexplorer.Version{Full: "^9.5"},
			},
		},
		Language: depexplorer.Language{
			Name: depexplorer.LanguageNamePHP,
			Version: &depexplorer.Version{
				Full: "8.1",
			},
		},
	}

	got, err := depexplorer.ExploreComposerJSON([]byte(file))
	require.NoError(t, err)
	require.Equal(t, expected, got)
}

func TestExploreComposerLock(t *testing.T) {
	file := `{
    "packages": [
        {
            "name": "phpunit/phpunit",
			"version": "7.1"
		}
	],
	"platform": {
		"php": "8.1"
	}
}`

	expected := &depexplorer.File{
		Name:              "composer.lock",
		Path:              "composer.lock",
		DependencyManager: depexplorer.DependencyManagerComposer,
		Dependencies: []*depexplorer.Dependency{
			{
				Name:    "phpunit/phpunit",
				Version: depexplorer.Version{Full: "7.1"},
			},
		},
		Language: depexplorer.Language{
			Name: depexplorer.LanguageNamePHP,
			Version: &depexplorer.Version{
				Full: "8.1",
			},
		},
	}

	got, err := depexplorer.ExploreComposerLock([]byte(file))
	require.NoError(t, err)
	require.Equal(t, expected, got)
}
