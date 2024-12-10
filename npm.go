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
		Dependencies:      make([]*Dependency, 0, len(definition.Dependencies)+len(definition.DevDependencies)),
		Language: Language{
			Name: LanguageNameJS,
		},
		Frameworks: make([]*Framework, 0),
	}

	for name, version := range definition.Dependencies {
		depFile.addDependency(name, version)
	}

	for name, version := range definition.DevDependencies {
		depFile.addDependency(name, version)
	}

	return depFile, nil
}

func ExplorePackageLockJSON(file []byte) (*File, error) {
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

	depFile := &File{
		Name:              "package-lock.json",
		Path:              "package-lock.json",
		DependencyManager: DependencyManagerNPM,
		Dependencies:      make([]*Dependency, 0, len(pkg.Dependencies)+len(pkg.DevDependencies)),
		Language: Language{
			Name: LanguageNameJS,
		},
		Frameworks: make([]*Framework, 0),
	}

	result := make([]*Dependency, 0, len(pkg.Dependencies)+len(pkg.DevDependencies))

	for name, version := range pkg.Dependencies {
		depFile.addDependency(name, version)
	}

	for name, version := range pkg.DevDependencies {
		depFile.addDependency(name, version)
	}

	return &File{
		Name:              "package-lock.json",
		Path:              "package-lock.json",
		DependencyManager: DependencyManagerNPM,
		Dependencies:      result,
		Language: Language{
			Name: LanguageNameJS,
		},
	}, nil
}
