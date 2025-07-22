package migrator

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

// Migrator содержит конфигурацию подключения к базе данных и параметры миграции.
type Migrator struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MigrationsPath  string
	MigrationsTable string
}

// dsn формирует строку подключения к PostgreSQL с учётом дополнительных параметров миграции.
func (r *Migrator) dsn() string {
	if r.SSLMode == "" {
		r.SSLMode = "disable"
	}
	if r.MigrationsTable == "" {
		r.MigrationsTable = "schema_migrations"
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&x-migrations-table=%s",
		r.User, r.Password, r.Host, r.Port, r.DBName, r.SSLMode, r.MigrationsTable,
	)
}

// runUp выполняет применение всех доступных миграций.
func (r *Migrator) runUp() error {
	rsd, err := migrate.New(
		"file://"+r.MigrationsPath,
		r.dsn(),
	)
	if err != nil {
		return fmt.Errorf("error creating migrator: %w", err)
	}
	defer rsd.Close()

	// Применяем миграции вверх
	if err := rsd.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	return nil
}

// runDown выполняет откат всех миграций.
func (r *Migrator) runDown() error {
	m, err := migrate.New(
		"file://"+r.MigrationsPath,
		r.dsn(),
	)
	if err != nil {
		return fmt.Errorf("migrator initialization error: %w", err)
	}
	defer m.Close()

	// Откатываем миграции
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration rollback error: %w", err)
	}

	return nil
}

// RunMigrate — точка входа для выполнения миграции. Вызывается из main или init-пакета.
func RunMigrate() {
	// Определение флагов командной строки
	mode := flag.String("mode", "", "migration mode: up or down (required)")
	dbHost := flag.String("db-host", "", "host postgres (required)")
	dbPort := flag.Int("db-port", 5432, "port postgres")
	dbUser := flag.String("db-user", "", "user postgres (required)")
	dbPass := flag.String("db-pass", "", "password postgres (required)")
	dbName := flag.String("db-name", "", "db name postgres (required)")
	sslMode := flag.String("ssl-mode", "disable", "ssl mode postgres")
	migrationsPath := flag.String("migrations-path", "", "path to migration (required)")
	migrationsTable := flag.String("migrations-table", "schema_migrations", "migration table name")

	flag.Parse()

	// Проверка обязательных параметров
	missing := []string{}
	if *mode == "" {
		missing = append(missing, "--mode")
	}
	if *dbHost == "" {
		missing = append(missing, "--db-host")
	}
	if *dbUser == "" {
		missing = append(missing, "--db-user")
	}
	if *dbName == "" {
		missing = append(missing, "--db-name")
	}
	if *dbPass == "" {
		missing = append(missing, "--db-pass")
	}
	if *migrationsPath == "" {
		missing = append(missing, "--migrations-path")
	}

	if len(missing) > 0 {
		fmt.Fprintf(os.Stderr, "❌  Mandatory flags not passed: %v\n", missing)
		flag.Usage()
		os.Exit(1)
	}

	// Создание мигратора на основе параметров
	r := &Migrator{
		Host:            *dbHost,
		Port:            *dbPort,
		User:            *dbUser,
		Password:        *dbPass,
		DBName:          *dbName,
		SSLMode:         *sslMode,
		MigrationsPath:  *migrationsPath,
		MigrationsTable: *migrationsTable,
	}

	// Выполнение миграции в зависимости от режима
	var err error
	switch *mode {
	case "up":
		err = r.runUp()
	case "down":
		err = r.runDown()
	default:
		fmt.Fprintf(os.Stderr, "❌ Unknown migration mode: %s. Use 'up' или 'down'\n", *mode)
		os.Exit(1)
	}

	// Обработка результата
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error while performing migration: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅  Migration completed successfully")
}
