# depexplorer

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/artarts36/depexplorer) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/artarts36/depexplorer/master/LICENSE)

```
go get github.com/artarts36/depexplorer
```

depexplorer - Go library for explore project dependencies from:
- go.mod
- composer.json
- composer.lock
- package.json
- package-lock.json

Result struct contains:
- Name of dependency manager
- List of dependencies and that versions
- List of used frameworks
- Name of programming language

Repository structure:
- `./`                    - main module with explore functions
- `./pkg/repository`      - module for explore in repositories
- `./pkg/github`          - module implements repository Client for GitHub
- `./pkg/gitlab`          - module implements repository Client for Gitlab
- `./pkg/repository-slog` - module implements slog adapter for repository logger

## Explore Go dependencies

```go
package main

import (
	"fmt"
	"github.com/artarts36/depexplorer"
	"os"
)

func main() {
	file, _ := os.ReadFile("/path/to/mod")
	depFile, _ := depexplorer.ExploreGoMod(file)
	for _, dep := range depFile.Dependencies {
		fmt.Println(dep.Name, dep.Version.Full)
	}
}
```

## Explore Go dependencies by path

```go
package main

import (
	"fmt"
	"github.com/artarts36/depexplorer"
)

func main() {
	depFile, _ := depexplorer.Explore("/path/to/mod", depexplorer.ExploreGoMod)
	for _, dep := range depFile.Dependencies {
		fmt.Println(dep.Name, dep.Version.Full)
	}
}
```

## Guess dependency manager and explore dependencies by path

```go
package main

import (
	"fmt"
	"github.com/artarts36/depexplorer"
)

func main() {
	file, _ := depexplorer.Guess("/path/to/package.json")
	fmt.Println(file.DependencyManager)
	for _, dep := range file.Dependencies {
		fmt.Println(dep.Name, dep.Version.Full)
	}
}
```

## Explore from GitHub Repository

You need to install another packages:
```
go get github.com/artarts36/depexplorer/pkg/repository
go get github.com/artarts36/depexplorer/pkg/github
```

And use this snippet:

```go
package main

import (
	"context"
	"fmt"
	"github.com/artarts36/depexplorer/pkg/github"
	"github.com/artarts36/depexplorer/pkg/repository"
	githubClient "github.com/google/go-github/github"
)

func main() {
	explorer := repository.NewExplorer(github.NewClient(nil))
	
	repo, _ := repository.NewRepoFromURI("https://github.com/artarts36/depexplorer")
	
	files, _ := explorer.ExploreRepository(context.Background(), repo, nil)
	for _, file := range files {
		fmt.Println(file.Name)
		
		for _, dep := range file.Dependencies {
			fmt.Println(dep.Name, dep.Version.Full)
        }
	}
}
```

## Explore from Gitlab Repository

You need to install another packages:
```
go get github.com/artarts36/depexplorer/pkg/repository
go get github.com/artarts36/depexplorer/pkg/gitlab
```

And use this snippet:

```go
package main

import (
	"context"
	"fmt"
	"github.com/artarts36/depexplorer/pkg/repository"
	"github.com/artarts36/depexplorer/pkg/gitlab"
)

func main() {
	client, _ := gitlab.NewClientWithToken("token", nil)
	
	explorer := repository.NewExplorer(client)
	
	repo, _ := repository.NewRepoFromURI("https://github.com/artarts36/depexplorer")
	
	files, _ := explorer.ExploreRepository(context.Background(), repo, nil)
	for _, file := range files {
		fmt.Println(file.Name)
		
		for _, dep := range file.Dependencies {
			fmt.Println(dep.Name, dep.Version.Full)
        }
	}
}
```

## Explore from GitHub and Gitlab Repositories

You need to install another packages:
```
go get github.com/artarts36/depexplorer/pkg/repository
go get github.com/artarts36/depexplorer/pkg/github
go get github.com/artarts36/depexplorer/pkg/gitlab
```

And use this snippet:

```go
package main

import (
	"context"
	"fmt"
	"github.com/artarts36/depexplorer/pkg/github"
	"github.com/artarts36/depexplorer/pkg/repository"
	"github.com/artarts36/depexplorer/pkg/gitlab"
)

func main() {
	clientsMap := map[string]repository.Client{
		"github.com": github.NewClient(nil),
	}

	gitlabClient, _ := gitlab.NewClientWithToken("token", nil)
	clientsMap["gitlab.com"] = gitlabClient

	explorer := repository.NewExplorer(repository.NewClientComposite(clientsMap))

	repo, _ := repository.NewRepoFromURI("https://github.com/artarts36/depexplorer")

	files, _ := explorer.ExploreRepository(context.Background(), repo, nil)
	for _, file := range files {
		fmt.Println(file.Name)

		for _, dep := range file.Dependencies {
			fmt.Println(dep.Name, dep.Version.Full)
		}
	}
}
```
