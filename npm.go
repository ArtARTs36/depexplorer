package depexplorer

import (
	"encoding/json"
	"fmt"
)

type packageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type packageLock struct {
	Packages map[string]packageJSON `json:"packages"`
}

func ExplorePackageJSON(file []byte) ([]*Dependency, error) {
	var definition packageJSON

	err := json.Unmarshal(file, &definition)
	if err != nil {
		return nil, fmt.Errorf("unable to parse npm definition: %w", err)
	}

	result := make([]*Dependency, 0, len(definition.Dependencies)+len(definition.DevDependencies))

	addDep := func(name, version string) {
		result = append(result, &Dependency{
			Name: name,
			Version: Version{
				Full: version,
			},
		})
	}

	for name, version := range definition.Dependencies {
		addDep(name, version)
	}

	for name, version := range definition.DevDependencies {
		addDep(name, version)
	}

	return result, nil
}

func ExplorePackageLockJSON(file []byte) ([]*Dependency, error) {
	var definition packageLock

	err := json.Unmarshal(file, &definition)
	if err != nil {
		return nil, fmt.Errorf("unable to parse npm definition: %w", err)
	}

	if len(definition.Packages) == 0 {
		return nil, fmt.Errorf("no packages found in npm definition")
	}

	pkg, ok := definition.Packages[""]
	if !ok {
		return nil, fmt.Errorf("no root package found in npm definition")
	}

	result := make([]*Dependency, 0, len(pkg.Dependencies)+len(pkg.DevDependencies))

	addDep := func(name, version string) {
		result = append(result, &Dependency{
			Name: name,
			Version: Version{
				Full: version,
			},
		})
	}

	for name, version := range pkg.Dependencies {
		addDep(name, version)
	}

	for name, version := range pkg.DevDependencies {
		addDep(name, version)
	}

	return result, nil
}
