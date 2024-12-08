package github

type Logger func(string, map[string]interface{})

func NoopLogger() Logger {
	return func(string, map[string]interface{}) {}
}
