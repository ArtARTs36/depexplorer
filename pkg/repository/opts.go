package repository

type ExploreOpts struct {
	Ref       string
	Directory string
}

func (o *ExploreOpts) getDirectory() string {
	if o == nil {
		return ""
	}
	return o.Directory
}

func (o *ExploreOpts) getRef() string {
	if o == nil {
		return ""
	}
	return o.Ref
}
