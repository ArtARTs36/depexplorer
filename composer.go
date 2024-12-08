package depexplorer

import (
	"encoding/json"
	"fmt"
)

type composerJSON struct {
	Require    map[string]string `json:"require"`
	RequireDev map[string]string `json:"require-dev"`
}

type composerLock struct {
	Packages []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"packages"`
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
	}, nil
}

func ExploreComposerLock(file []byte) (*File, error) {
	var definition composerLock

	err := json.Unmarshal(file, &definition)
	if err != nil {
		return nil, fmt.Errorf("unable to parse composer definition: %w", err)
	}

	if len(definition.Packages) < 1 {
		return &File{
			Name:              "composer.lock",
			Path:              "composer.lock",
			DependencyManager: DependencyManagerComposer,
			Dependencies:      make([]*Dependency, 0),
		}, nil
	}

	result := make([]*Dependency, 0, len(definition.Packages)-1)
	for i := 1; i < len(definition.Packages); i++ {
		pkg := definition.Packages[i]

		result = append(result, &Dependency{
			Name: pkg.Name,
			Version: Version{
				Full: pkg.Version,
			},
		})
	}

	return &File{
		Name:              "composer.lock",
		Path:              "composer.lock",
		DependencyManager: DependencyManagerComposer,
		Dependencies:      result,
	}, nil
}
