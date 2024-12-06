package depexplorer

type DependencyManager string

const (
	DependencyManagerNone     DependencyManager = ""
	DependencyManagerGoMod    DependencyManager = "go.mod"
	DependencyManagerComposer DependencyManager = "composer"
	DependencyManagerNPM      DependencyManager = "npm"
)

type Dependency struct {
	Name    string
	Version Version
}

type Version struct {
	Full string
}
