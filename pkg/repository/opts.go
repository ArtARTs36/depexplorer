package repository

type ExploreOpts struct {
	Directory string
}

func (o *ExploreOpts) getDirectory() string {
	if o == nil {
		return ""
	}
	return o.Directory
}
