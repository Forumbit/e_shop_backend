package main

import (
	"e_shop_backend.amirkharisov.net/internal/infrastructure/config"
	"e_shop_backend.amirkharisov.net/migrations"
	"e_shop_backend.amirkharisov.net/pkg/postgre"
)

func main() {
	cfg := config.MustLoad()

	db, err := postgre.NewDBConfig(cfg.DB)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	defer db.Close()

	if err = migrations.PostgresMigrate(db, migrations.PostgresEmbedFS, cfg.DB); err != nil {
		panic("Failed to run migrations: " + err.Error())
	}
}
