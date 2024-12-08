package depexplorer

import (
	"fmt"
	"path/filepath"
)

type guessedFile struct {
	Path              string
	Name              string
	DependencyManager DependencyManager
	IsLockFile        bool
	CanHaveLockFile   bool
	Explorer          FileExplorer
}

func Guess(path string) (*File, error) {
	guessed, err := guess(path)
	if err != nil {
		return nil, err
	}

	return exploreGuessedFile(guessed, Explore)
}

func guess(path string) (*guessedFile, error) {
	guessed := &guessedFile{
		Path:              path,
		Name:              filepath.Base(path),
		DependencyManager: DependencyManagerNone,
	}

	switch guessed.Name {
	case "composer.lock":
		guessed.DependencyManager = DependencyManagerComposer
		guessed.Explorer = ExploreComposerLock
		guessed.IsLockFile = true
	case "composer.json":
		guessed.DependencyManager = DependencyManagerComposer
		guessed.Explorer = ExploreComposerJSON
		guessed.CanHaveLockFile = true
	case "go.mod":
		guessed.DependencyManager = DependencyManagerGoMod
		guessed.Explorer = ExploreGoMod
	case "package.json":
		guessed.DependencyManager = DependencyManagerNPM
		guessed.Explorer = ExplorePackageJSON
		guessed.CanHaveLockFile = true
	case "package-lock.json":
		guessed.DependencyManager = DependencyManagerNPM
		guessed.Explorer = ExplorePackageLockJSON
		guessed.IsLockFile = true
	}

	if guessed.Explorer == nil {
		return guessed, fmt.Errorf("could not guess dependency from %s", path)
	}

	return guessed, nil
}

func exploreGuessedFile(guessed *guessedFile, contentExplorer fileContentExplorer) (*File, error) {
	deps, err := contentExplorer(guessed.Path, guessed.Explorer)
	if err != nil {
		return &File{
			Name:              guessed.Name,
			Path:              guessed.Path,
			DependencyManager: guessed.DependencyManager,
		}, err
	}

	return &File{
		Name:              guessed.Name,
		Path:              guessed.Path,
		DependencyManager: guessed.DependencyManager,
		Dependencies:      deps.Dependencies,
	}, nil
}
