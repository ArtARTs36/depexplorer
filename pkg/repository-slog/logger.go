package repository_slog

import (
	"fmt"
	"log/slog"

	"github.com/artarts36/depexplorer/pkg/repository"
)

func New() repository.Logger {
	return Prefix("")
}

func Prefix(prefix string) repository.Logger {
	return func(s string, m map[string]interface{}) {
		l := slog.Default()

		for k, v := range m {
			l = l.With(slog.Any(k, v))
		}

		l.Info(fmt.Sprintf("%s%s", prefix, s))
	}
}
