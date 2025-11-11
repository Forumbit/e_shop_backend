package services

import (
	"context"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
)

type brandRepository interface {
	Create(ctx context.Context, brand *models.Brand) (int, error)
	GetByID(ctx context.Context, id int) (*models.Brand, error)
	Update(ctx context.Context, brand *models.Brand) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Brand, error)
}

type BrandService struct {
	repo brandRepository
}

func NewBrandService(brandRepository brandRepository) *BrandService {
	return &BrandService{repo: brandRepository}
}

func (bs *BrandService) Create(ctx context.Context, brand *models.Brand) (int, error) {
	return bs.repo.Create(ctx, brand)
}

func (bs *BrandService) GetByID(ctx context.Context, id int) (*models.Brand, error) {
	return bs.repo.GetByID(ctx, id)
}

func (bs *BrandService) Update(ctx context.Context, brand *models.Brand) error {
	return bs.repo.Update(ctx, brand)
}

func (bs *BrandService) Delete(ctx context.Context, id int) error {
	return bs.repo.Delete(ctx, id)
}

func (bs *BrandService) GetAll(ctx context.Context) ([]models.Brand, error) {
	return bs.repo.GetAll(ctx)
}
