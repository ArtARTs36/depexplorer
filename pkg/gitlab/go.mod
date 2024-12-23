module github.com/artarts36/depexplorer/pkg/gitlab

go 1.23.3

replace (
	github.com/artarts36/depexplorer => ./../..
	github.com/artarts36/depexplorer/pkg/repository => ./../repository
)

require (
	github.com/artarts36/depexplorer/pkg/repository v0.0.0-00010101000000-000000000000
	gitlab.com/gitlab-org/api/client-go v0.116.0
)

require (
	github.com/artarts36/depexplorer v0.1.0 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/oauth2 v0.6.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.29.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
