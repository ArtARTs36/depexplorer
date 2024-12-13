package repository

type Repo struct {
	Owner string // required
	Name  string // required

	Directory string // optional
}
