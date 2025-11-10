package services

import (
	"context"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
)

type productRepository interface {
	Create(ctx context.Context, product *models.Product) (int, error)
	GetByID(ctx context.Context, id int) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Product, error)
}

type ProductService struct {
	productRepo productRepository
}

func NewProductService(productRepo productRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) Create(ctx context.Context, product *models.Product) (int, error) {
	return s.productRepo.Create(ctx, product)
}

func (s *ProductService) GetByID(ctx context.Context, id int) (*models.Product, error) {
	return s.productRepo.GetByID(ctx, id)
}
func (s *ProductService) Update(ctx context.Context, product *models.Product) error {
	return s.productRepo.Update(ctx, product)
}
func (s *ProductService) Delete(ctx context.Context, id int) error {
	return s.productRepo.Delete(ctx, id)
}
func (s *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	return s.productRepo.GetAll(ctx)
}
