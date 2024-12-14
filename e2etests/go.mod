module e2etests

go 1.23.3

replace (
	github.com/artarts36/depexplorer/pkg/github => ./../pkg/github
	github.com/artarts36/depexplorer/pkg/repository => ./../pkg/repository
	github.com/artarts36/depexplorer/pkg/repository-slog => ./../pkg/repository-slog
)

require (
	github.com/artarts36/depexplorer v0.1.0
	github.com/artarts36/depexplorer/pkg/github v0.0.0-00010101000000-000000000000
	github.com/artarts36/depexplorer/pkg/repository v0.1.0
	github.com/artarts36/depexplorer/pkg/repository-slog v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-github/v67 v67.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	golang.org/x/mod v0.22.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
