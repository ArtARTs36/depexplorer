package depexplorer

import (
	"fmt"
	"path/filepath"
)

func Guess(path string) (DependencyManager, []*Dependency, error) {
	var explorer FileExplorer

	filename := filepath.Base(path)
	depManager := DependencyManagerNone

	switch filename {
	case "composer.json":
		explorer = ExploreComposerJSON
		depManager = DependencyManagerComposer
	case "composer.lock":
		explorer = ExploreComposerLock
		depManager = DependencyManagerComposer
	case "go.mod":
		explorer = ExploreGoMod
		depManager = DependencyManagerGoMod
	case "package.json":
		explorer = ExplorePackageJSON
		depManager = DependencyManagerNPM
	case "package-lock.json":
		explorer = ExplorePackageLockJSON
		depManager = DependencyManagerNPM
	}

	if explorer == nil {
		return DependencyManagerNone, nil, fmt.Errorf("could not guess dependency from %s", path)
	}

	deps, err := Explore(path, explorer)
	if err != nil {
		return depManager, nil, err
	}

	return depManager, deps, nil
}
