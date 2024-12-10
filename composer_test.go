package depexplorer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/artarts36/depexplorer"
)

func TestExploreComposerJSON(t *testing.T) {
	file := `{
    "name": "artarts36/local-file-system",
    "require": {
		"php": "8.1",
		"symfony/framework-bundle": "5.0.0"
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
				Name:    "symfony/framework-bundle",
				Version: depexplorer.Version{Full: "5.0.0"},
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
		Frameworks: []*depexplorer.Framework{
			{
				Name: depexplorer.FrameworkNameSymfony,
				Version: depexplorer.Version{
					Full: "5.0.0",
				},
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
