package repository

import (
	"fmt"
	"net/url"
	"strings"
)

type Repo struct {
	Host string // optional

	Owner string // required
	Name  string // required
}

func NewGithubRepo(owner string, name string) Repo {
	return Repo{
		Host:  "github.com",
		Owner: owner,
		Name:  name,
	}
}

func NewRepoFromURI(uri string) (Repo, error) {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return Repo{}, fmt.Errorf("failed to parse uri: %w", err)
	}

	pathParts := strings.Split(strings.Trim(parsedURI.Path, "/"), "/")
	if len(pathParts) == 0 {
		return Repo{}, fmt.Errorf("failed to parse repo uri: path parts is empty")
	}

	return Repo{
		Host:  parsedURI.Host,
		Owner: strings.Join(pathParts[0:len(pathParts)-1], "/"),
		Name:  pathParts[len(pathParts)-1],
	}, nil
}
