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

	depFile := &File{
		Name:              "go.mod",
		Path:              "go.mod",
		DependencyManager: DependencyManagerGoMod,
		Dependencies:      make([]*Dependency, 0, len(mod.Require)),
		Language: Language{
			Name: LanguageNameGo,
		},
		Frameworks: make([]*Framework, 0),
	}

	for _, require := range mod.Require {
		depFile.addDependency(require.Mod.Path, require.Mod.Version)
	}

	if mod.Go != nil {
		depFile.Language.Version = &Version{
			Full: mod.Go.Version,
		}
	}

	return depFile, nil
}
