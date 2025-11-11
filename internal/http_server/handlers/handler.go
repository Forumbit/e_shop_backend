package handlers

import (
	"context"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
)

type categoryService interface {
	Create(ctx context.Context, category *models.Category) (int, error)
	GetByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Category, error)
}

type productService interface {
	Create(ctx context.Context, product *models.Product) (int, error)
	GetByID(ctx context.Context, id int) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Product, error)
}

type Handler struct {
	ctg  categoryService
	prrs productService
}

func NewHandler(categoryService categoryService, productService productService) *Handler {
	return &Handler{
		prrs: productService,
		ctg:  categoryService,
	}
}
