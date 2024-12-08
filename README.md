# depexplorer

depexplorer - Go library for explore project dependencies from:
- go.mod
- composer.json
- composer.lock
- package.json
- package-lock.json

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
	"os"
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
	"os"
)

func main() {
	file, _ := depexplorer.Guess("/path/to/package.json")
	fmt.Println(file.DependencyManager)
	for _, dep := range file.Dependencies {
		fmt.Println(dep.Name, dep.Version.Full)
	}
}
```
