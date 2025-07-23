package app

import (
	"github.com/Jonatna0990/test-subscription-service/pkg/postgres"
	"log"
)

// initPostgres инициализирует подключение к PostgreSQL
func (a *App) initPostgres() {
	psg, err := postgres.NewPostgres(&a.config.Postgres)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	a.logger.Info("Postgres initialized")
	a.postgres = psg
}
