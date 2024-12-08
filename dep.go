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

	Language Language
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
