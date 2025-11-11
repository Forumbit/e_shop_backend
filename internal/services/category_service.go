package services

import (
	"context"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
)

type categoryRepository interface {
	Create(ctx context.Context, category *models.Category) (int, error)
	GetByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Category, error)
}

type CategoryService struct {
	repo categoryRepository
}

func NewCategoryService(repo categoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (cs *CategoryService) Create(ctx context.Context, category *models.Category) (int, error) {
	return cs.repo.Create(ctx, category)
}

func (cs *CategoryService) GetByID(ctx context.Context, id int) (*models.Category, error) {
	return cs.repo.GetByID(ctx, id)
}

func (cs *CategoryService) Update(ctx context.Context, category *models.Category) error {
	return cs.repo.Update(ctx, category)
}

func (cs *CategoryService) Delete(ctx context.Context, id int) error {
	return cs.repo.Delete(ctx, id)
}

func (cs *CategoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	return cs.repo.GetAll(ctx)
}
