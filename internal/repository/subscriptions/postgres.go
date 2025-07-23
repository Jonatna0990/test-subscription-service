package subscriptions

import (
	"github.com/Jonatna0990/test-subscription-service/pkg/postgres"
	"log/slog"
)

type repo struct {
	logger   *slog.Logger
	postgres *postgres.Postgres
}

func NewRepository(logger *slog.Logger, postgres *postgres.Postgres) Repository {
	return &repo{logger: logger, postgres: postgres}
}
