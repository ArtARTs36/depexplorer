package github

type Logger interface {
	Printf(string, map[string]interface{})
}

type NoopLogger struct{}

func (NoopLogger) Printf(string, map[string]interface{}) {}
