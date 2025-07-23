package migrator

import (
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/pkg/migrator"
	"github.com/spf13/cobra"
	"os"
)

func RunMigrate() *cobra.Command {
	var dbHost, dbUser, dbPass, dbName, sslMode, mode, migrationsPath, migrationsTable string
	var dbPort int

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run migrations",
		Run: func(cmd *cobra.Command, args []string) {
			err := migrator.RunMigrate(dbHost, dbPort, dbUser, dbPass, dbName, sslMode, mode, migrationsPath, migrationsTable)
			if err != nil {
				fmt.Errorf("migration failed: %w", err)
			}
			// TODO
			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&dbHost, "db-host", "", "DB host")
	cmd.Flags().IntVar(&dbPort, "db-port", 5432, "DB port")
	cmd.Flags().StringVar(&dbUser, "db-user", "", "DB user")
	cmd.Flags().StringVar(&dbPass, "db-pass", "", "DB password")
	cmd.Flags().StringVar(&dbName, "db-name", "", "DB name")
	cmd.Flags().StringVar(&sslMode, "ssl-mode", "disable", "DB SSL mode")
	cmd.Flags().StringVar(&mode, "mode", "", "Migration mode: up or down")
	cmd.Flags().StringVar(&migrationsPath, "migrations-path", "", "Path to migrations")
	cmd.Flags().StringVar(&migrationsTable, "migrations-table", "schema_migrations", "Migrations table name")

	_ = cmd.MarkFlagRequired("mode")
	_ = cmd.MarkFlagRequired("db-host")
	_ = cmd.MarkFlagRequired("db-user")
	_ = cmd.MarkFlagRequired("db-pass")
	_ = cmd.MarkFlagRequired("db-name")
	_ = cmd.MarkFlagRequired("migrations-path")

	return cmd
}
