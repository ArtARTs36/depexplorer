package depexplorer

import (
	"encoding/json"
	"fmt"
)

const composerPHPDependency = "php"

type composerJSON struct {
	Require    map[string]string `json:"require"`
	RequireDev map[string]string `json:"require-dev"`

	Config struct {
		Platform map[string]string `json:"platform"`
	} `json:"config"`
}

type composerLock struct {
	Packages []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"packages"`
	Platform map[string]string `json:"platform"`
}

func findPHPVersionInComposerJSON(definition composerJSON) *Version {
	v := definition.Config.Platform[composerPHPDependency]
	if v != "" {
		return &Version{
			Full: v,
		}
	}

	v = definition.Require[composerPHPDependency]
	if v != "" {
		return &Version{
			Full: v,
		}
	}

	v = definition.RequireDev[composerPHPDependency]
	if v != "" {
		return &Version{
			Full: v,
		}
	}

	return nil
}

func ExploreComposerJSON(file []byte) (*File, error) {
	var definition composerJSON

	err := json.Unmarshal(file, &definition)
	if err != nil {
		return nil, fmt.Errorf("unable to parse composer definition: %w", err)
	}

	result := make([]*Dependency, 0, len(definition.Require)+len(definition.RequireDev))

	addDep := func(name, version string) {
		result = append(result, &Dependency{
			Name: name,
			Version: Version{
				Full: version,
			},
		})
	}

	for name, version := range definition.Require {
		addDep(name, version)
	}

	for name, version := range definition.RequireDev {
		addDep(name, version)
	}

	return &File{
		Name:              "composer.json",
		Path:              "composer.json",
		DependencyManager: DependencyManagerComposer,
		Dependencies:      result,
		Language: Language{
			Name:    LanguageNamePHP,
			Version: findPHPVersionInComposerJSON(definition),
		},
	}, nil
}

func ExploreComposerLock(file []byte) (*File, error) {
	var definition composerLock

	err := json.Unmarshal(file, &definition)
	if err != nil {
		return nil, fmt.Errorf("unable to parse composer definition: %w", err)
	}

	result := make([]*Dependency, 0, len(definition.Packages))
	for i := 0; i < len(definition.Packages); i++ {
		pkg := definition.Packages[i]

		result = append(result, &Dependency{
			Name: pkg.Name,
			Version: Version{
				Full: pkg.Version,
			},
		})
	}

	var phpVersion *Version
	if pVersion := definition.Platform["php"]; pVersion != "" {
		phpVersion = &Version{
			Full: pVersion,
		}
	}

	return &File{
		Name:              "composer.lock",
		Path:              "composer.lock",
		DependencyManager: DependencyManagerComposer,
		Dependencies:      result,
		Language: Language{
			Name:    LanguageNamePHP,
			Version: phpVersion,
		},
	}, nil
}
