package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (cr *CategoryRepository) Create(ctx context.Context, category *models.Category) (int, error) {
	stmt, err := cr.db.PrepareContext(ctx, "INSERT INTO categories (name) VALUES ($1) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	var id int

	if err = stmt.QueryRowContext(ctx, category.Name).Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err.Error())
	}

	return id, nil
}

func (cr *CategoryRepository) GetByID(ctx context.Context, id int) (*models.Category, error) {
	stmt, err := cr.db.PrepareContext(ctx, "SELECT * FROM categories WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	var category models.Category

	if err = stmt.QueryRowContext(ctx, id).Scan(&category.ID, &category.Name); err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err.Error())
	}
	return &category, nil
}

func (cr *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	stmt, err := cr.db.PrepareContext(ctx, "UPDATE categories SET name = $2 WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, &category.ID, &category.Name); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err.Error())
	}

	return nil
}

func (cr *CategoryRepository) Delete(ctx context.Context, id int) error {
	stmt, err := cr.db.PrepareContext(ctx, "DELETE FROM categories WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err.Error())
	}

	return nil
}

func (cr *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	stmt, err := cr.db.PrepareContext(ctx, "SELECT * FROM categories LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	var categories []models.Category

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		rows.Scan(&category.ID, &category.Name)
		categories = append(categories, category)
	}
	return categories, nil
}
