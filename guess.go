package depexplorer

import (
	"fmt"
	"path/filepath"
)

func Guess(path string) ([]*Dependency, error) {
	var explorer FileExplorer

	filename := filepath.Base(path)

	switch filename {
	case "composer.json":
		explorer = ExploreComposerJSON
	case "composer.lock":
		explorer = ExploreComposerLock
	case "go.mod":
		explorer = ExploreGoMod
	case "package.json":
		explorer = ExplorePackageJSON
	case "package-lock.json":
		explorer = ExplorePackageLockJSON
	}

	if explorer == nil {
		return nil, fmt.Errorf("could not guess dependency from %s", path)
	}

	return Explore(path, explorer)
}
