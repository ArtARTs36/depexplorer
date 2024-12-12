# depexplorer

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/artarts36/depexplorer) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/artarts36/depexplorer/master/LICENSE)

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

Install as: `go get github.com/artarts36/depexplorer`

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

You need to install another package:
```
go get github.com/artarts36/depexplorer/pkg/github
```

And use this snippet:

```go
package main

import (
	"context"
	"fmt"
	"github.com/artarts36/depexplorer/pkg/github"
	githubClient "github.com/google/go-github/github"
)

func main() {
	ghClient := githubClient.NewClient(nil)

	file, _ := github.ScanRepository(context.Background(), ghClient, github.Repository{
		Owner: "artarts36",
		Repo:  "depexplorer",
	})
	fmt.Println(file.DependencyManager)
	for _, dep := range file.Dependencies {
		fmt.Println(dep.Name, dep.Version.Full)
	}
}
```
