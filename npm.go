package depexplorer

import (
	"encoding/json"
	"fmt"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type packageJSON struct {
	Dependencies    orderedmap.OrderedMap[string, string] `json:"dependencies"`
	DevDependencies orderedmap.OrderedMap[string, string] `json:"devDependencies"`
}

type packageLock struct {
	Packages orderedmap.OrderedMap[string, packageJSON] `json:"packages"`
}

func ExplorePackageJSON(file []byte) (*File, error) {
	var definition packageJSON

	err := json.Unmarshal(file, &definition)
	if err != nil {
		return nil, fmt.Errorf("unable to parse npm definition: %w", err)
	}

	depFile := &File{
		Name:              "package.json",
		Path:              "package.json",
		DependencyManager: DependencyManagerNPM,
		Dependencies:      make([]*Dependency, 0, definition.Dependencies.Len()+definition.DevDependencies.Len()),
		Language: Language{
			Name: LanguageNameJS,
		},
		Frameworks: make([]*Framework, 0),
	}

	depFile.addDependenciesFromOrderedMap(definition.Dependencies)
	depFile.addDependenciesFromOrderedMap(definition.DevDependencies)

	return depFile, nil
}

func ExplorePackageLockJSON(file []byte) (*File, error) {
	var definition packageLock

	err := json.Unmarshal(file, &definition)
	if err != nil {
		return nil, fmt.Errorf("unable to parse npm definition: %w", err)
	}

	if definition.Packages.Len() == 0 {
		return nil, fmt.Errorf("no packages found in npm definition")
	}

	pkg, ok := definition.Packages.Get("")
	if !ok {
		return nil, fmt.Errorf("no root package found in npm definition")
	}

	depFile := &File{
		Name:              "package-lock.json",
		Path:              "package-lock.json",
		DependencyManager: DependencyManagerNPM,
		Dependencies:      make([]*Dependency, 0, pkg.Dependencies.Len()+pkg.DevDependencies.Len()),
		Language: Language{
			Name: LanguageNameJS,
		},
		Frameworks: make([]*Framework, 0),
	}

	depFile.addDependenciesFromOrderedMap(pkg.Dependencies)
	depFile.addDependenciesFromOrderedMap(pkg.DevDependencies)

	return depFile, nil
}
