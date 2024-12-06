package depexplorer

import (
	"fmt"
	"golang.org/x/mod/modfile"
)

func ExploreGoMod(file []byte) ([]*Dependency, error) {
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

	return deps, nil
}
