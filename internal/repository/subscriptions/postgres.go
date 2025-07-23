package subscriptions

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"log/slog"
)

type DB interface {
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
}

type repo struct {
	logger *slog.Logger
	db     DB
}

func NewRepository(db DB, logger *slog.Logger) Repository {
	return &repo{db: db, logger: logger}
}
