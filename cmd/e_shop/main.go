package main

import (
	"e_shop_backend.amirkharisov.net/internal/http_server"
	"e_shop_backend.amirkharisov.net/internal/http_server/handlers"
	"e_shop_backend.amirkharisov.net/internal/infrastructure/config"
	"e_shop_backend.amirkharisov.net/internal/repositories"
	"e_shop_backend.amirkharisov.net/internal/services"
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

	categoryRepo := repositories.NewCategoryRepository(db)
	brandRepo := repositories.NewBrandRepository(db)
	productRepo := repositories.NewProductRepository(db)

	categoryService := services.NewCategoryService(categoryRepo)
	brandService := services.NewBrandService(brandRepo)
	productService := services.NewProductService(productRepo)

	handler := handlers.NewHandler(categoryService, productService, brandService)

	router := http_server.NewRouter(cfg, handler)

	router.Run()
}
