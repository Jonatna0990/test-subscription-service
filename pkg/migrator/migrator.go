package migrator

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// Migrator — конфигурация для выполнения миграций
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

// dsn возвращает строку подключения с параметрами миграций
func (r *Migrator) dsn() string {
	sslMode := r.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}
	table := r.MigrationsTable
	if table == "" {
		table = "schema_migrations"
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&x-migrations-table=%s",
		r.User, r.Password, r.Host, r.Port, r.DBName, sslMode, table,
	)
}

// run запускает миграции с указанным действием: "up" или "down"
func (r *Migrator) run(direction string) error {
	m, err := migrate.New("file://"+r.MigrationsPath, r.dsn())
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}
	defer m.Close()

	var opErr error
	switch direction {
	case "up":
		opErr = m.Up()
	case "down":
		opErr = m.Down()
	default:
		return fmt.Errorf("unknown direction: %s", direction)
	}

	if opErr != nil {
		if errors.Is(opErr, migrate.ErrNoChange) {
			fmt.Println("⚠️ no migrations to apply")
			return nil
		}
		return fmt.Errorf("migration %s failed: %w", direction, opErr)
	}

	return nil
}

// RunMigrate — точка входа. Запускает миграции на основе флагов
func RunMigrate(dbHost string, dbPort int, dbUser, dbPass, dbName, sslMode, mode, migrationsPath, migrationsTable string) error {
	// Проверка обязательных флагов
	required := map[string]string{
		"mode":            mode,
		"db-host":         dbHost,
		"db-user":         dbUser,
		"db-pass":         dbPass,
		"db-name":         dbName,
		"migrations-path": migrationsPath,
	}
	var missing []string
	for name, value := range required {
		if value == "" {
			missing = append(missing, "--"+name)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required flags: %v", missing)
	}

	migrator := Migrator{
		Host:            dbHost,
		Port:            dbPort,
		User:            dbUser,
		Password:        dbPass,
		DBName:          dbName,
		SSLMode:         sslMode,
		MigrationsPath:  migrationsPath,
		MigrationsTable: migrationsTable,
	}

	if err := migrator.run(mode); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("✅ Migration completed successfully")
	return nil
}
