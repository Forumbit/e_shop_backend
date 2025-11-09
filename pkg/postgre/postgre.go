package postgre

import (
	"database/sql"
	"fmt"

	"e_shop_backend.amirkharisov.net/internal/infrastructure/config"
)

func NewDBConfig(cfg config.DBconfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=public",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBname, cfg.SSLmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
