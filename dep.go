package depexplorer

type Dependency struct {
	Name    string
	Version Version
}

type Version struct {
	Full string
}
