package e2etests

import (
	"log/slog"

	"github.com/artarts36/depexplorer/pkg/repository"
)

func NewRepositoryExplorerLogger() repository.Logger {
	return func(s string, m map[string]interface{}) {
		l := slog.Default()

		for k, v := range m {
			l = l.With(slog.Any(k, v))
		}

		l.Info(s)
	}
}
