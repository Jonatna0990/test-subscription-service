package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"time"
)

type Postgres struct {
	cfg  *Config
	Pool *pgxpool.Pool
}

// NewPostgres создаёт и инициализирует новое подключение к PostgreSQL.
func NewPostgres(cfg *Config) (*Postgres, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	//TODO Вынести параметры в конфигурацию

	// Пример настройки таймаутов
	config.ConnConfig.ConnectTimeout = 5 * time.Second            // для TCP-соединения (работает не во всех случаях)
	config.ConnConfig.RuntimeParams["statement_timeout"] = "5000" // PostgreSQL-level query timeout (в миллисекундах)

	// Настройки пула
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 30 * time.Second
	config.MaxConnLifetime = 5 * time.Minute

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		cfg:  cfg,
		Pool: pool,
	}, nil
}
