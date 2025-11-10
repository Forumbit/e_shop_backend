package handlers

import (
	"context"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
)

type productService interface {
	Create(ctx context.Context, product *models.Product) (int, error)
	GetByID(ctx context.Context, id int) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Product, error)
}

type Handler struct {
	prrs productService
}

func NewHandler(productService productService) *Handler {
	return &Handler{
		prrs: productService,
	}
}
