package depexplorer

import (
	"fmt"
	"golang.org/x/mod/modfile"
)

func ExploreGoMod(file []byte) (*File, error) {
	mod, err := modfile.ParseLax("go.mod", file, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go.mod: %w", err)
	}

	deps := make([]*Dependency, len(mod.Require))
	for i, require := range mod.Require {
		deps[i] = &Dependency{
			Name: require.Mod.Path,
			Version: Version{
				Full: require.Mod.Version,
			},
		}
	}

	return &File{
		Name:              "go.mod",
		Path:              "go.mod",
		DependencyManager: DependencyManagerGoMod,
		Dependencies:      deps,
		LanguageVersion: &Version{
			Full: mod.Go.Version,
		},
	}, nil
}
