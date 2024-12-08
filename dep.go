package depexplorer

type DependencyManager string

const (
	DependencyManagerNone     DependencyManager = ""
	DependencyManagerGoMod    DependencyManager = "go.mod"
	DependencyManagerComposer DependencyManager = "composer"
	DependencyManagerNPM      DependencyManager = "npm"
)

type File struct {
	Path string
	Name string

	DependencyManager DependencyManager
	Dependencies      []*Dependency

	LanguageVersion *Version
}

type Dependency struct {
	Name    string
	Version Version
}

type Version struct {
	Full string
}
