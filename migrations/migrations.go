package migrations

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"

	"e_shop_backend.amirkharisov.net/internal/infrastructure/config"
	"e_shop_backend.amirkharisov.net/internal/infrastructure/constants"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed postgres-migs
var PostgresEmbedFS embed.FS

func PostgresMigrate(sql *sql.DB, migFiles fs.FS, cfg config.DBconfig) error {
	goose.SetBaseFS(migFiles)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("Driver's fault: %w", err)
	}

	if err := goose.Up(sql, constants.MigrationsPath); err != nil {
		return fmt.Errorf("Failed migrations executing: %w", err)
	}

	return nil
}
