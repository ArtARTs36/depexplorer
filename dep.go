package depexplorer

type (
	DependencyManager string
	LanguageName      string
)

const (
	DependencyManagerNone     DependencyManager = ""
	DependencyManagerGoMod    DependencyManager = "go.mod"
	DependencyManagerComposer DependencyManager = "composer"
	DependencyManagerNPM      DependencyManager = "npm"

	LanguageNameNone LanguageName = "none"
	LanguageNameGo   LanguageName = "go"
	LanguageNamePHP  LanguageName = "php"
	LanguageNameJS   LanguageName = "js"
)

type File struct {
	Path string
	Name string

	DependencyManager DependencyManager
	Dependencies      []*Dependency

	Language   Language
	Frameworks []*Framework
}

type Language struct {
	Name    LanguageName
	Version *Version
}

type Dependency struct {
	Name    string
	Version Version
}

type Version struct {
	Full string
}

func (f *File) addDependency(name string, version string) {
	dependency := &Dependency{
		Name:    name,
		Version: Version{Full: version},
	}
	f.Dependencies = append(f.Dependencies, dependency)

	framework, isFramework := dependencyToFramework(f.DependencyManager, dependency)
	if isFramework {
		f.Frameworks = append(f.Frameworks, framework)
	}
}
