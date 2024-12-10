package depexplorer

import (
	"encoding/json"
	"fmt"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

const composerPHPDependency = "php"

type composerJSON struct {
	Require    orderedmap.OrderedMap[string, string] `json:"require"`
	RequireDev orderedmap.OrderedMap[string, string] `json:"require-dev"`

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

	v, ok := definition.Require.Get(composerPHPDependency)
	if ok {
		return &Version{
			Full: v,
		}
	}

	v, ok = definition.RequireDev.Get(composerPHPDependency)
	if ok {
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

	depFile := &File{
		Name:              "composer.json",
		Path:              "composer.json",
		DependencyManager: DependencyManagerComposer,
		Dependencies:      make([]*Dependency, 0, definition.Require.Len()+definition.RequireDev.Len()),
		Language: Language{
			Name:    LanguageNamePHP,
			Version: findPHPVersionInComposerJSON(definition),
		},
		Frameworks: make([]*Framework, 0),
	}

	for pair := definition.Require.Oldest(); pair != nil; pair = pair.Next() {
		depFile.addDependency(pair.Key, pair.Value)
	}

	for pair := definition.RequireDev.Oldest(); pair != nil; pair = pair.Next() {
		depFile.addDependency(pair.Key, pair.Value)
	}

	return depFile, nil
}

func ExploreComposerLock(file []byte) (*File, error) {
	var definition composerLock

	err := json.Unmarshal(file, &definition)
	if err != nil {
		return nil, fmt.Errorf("unable to parse composer definition: %w", err)
	}

	depFile := &File{
		Name:              "composer.lock",
		Path:              "composer.lock",
		DependencyManager: DependencyManagerComposer,
		Dependencies:      make([]*Dependency, 0, len(definition.Packages)),
		Language: Language{
			Name: LanguageNamePHP,
		},
	}

	for _, pkg := range definition.Packages {
		depFile.addDependency(pkg.Name, pkg.Version)
	}

	var phpVersion *Version
	if pVersion := definition.Platform["php"]; pVersion != "" {
		phpVersion = &Version{
			Full: pVersion,
		}
	}

	depFile.Language.Version = phpVersion

	return depFile, nil
}
