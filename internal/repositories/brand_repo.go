package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
)

type BrandRepository struct {
	db *sql.DB
}

func NewBrandRepository(db *sql.DB) *BrandRepository {
	return &BrandRepository{db: db}
}

func (br *BrandRepository) Create(ctx context.Context, brand *models.Brand) (int, error) {
	stmt, err := br.db.PrepareContext(ctx, "INSERT INTO brands (name) VALUES ($1) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statemtn: %w", err.Error())
	}
	defer stmt.Close()

	var id int

	if err = stmt.QueryRowContext(ctx, &brand.Name).Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err.Error())
	}
	return id, nil
}

func (br *BrandRepository) GetByID(ctx context.Context, id int) (*models.Brand, error) {
	stmt, err := br.db.PrepareContext(ctx, "SELECT * FROM brands WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	var brand models.Brand
	if err = stmt.QueryRow(id).Scan(&brand.ID, &brand.Name); err != nil {
		return nil, fmt.Errorf("failed to execute statement", err.Error())
	}
	return &brand, nil
}

func (br *BrandRepository) Update(ctx context.Context, brand *models.Brand) error {
	stmt, err := br.db.PrepareContext(ctx, "UPDATE brands SET name = $2 WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, brand.ID, brand.Name); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err.Error())
	}

	return nil
}

func (br *BrandRepository) Delete(ctx context.Context, id int) error {
	stmt, err := br.db.PrepareContext(ctx, "DELETE FROM brands WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("failed to exectue statement: %w", err.Error())
	}
	return nil
}

func (br *BrandRepository) GetAll(ctx context.Context) ([]models.Brand, error) {
	stmt, err := br.db.PrepareContext(ctx, "SELECT * FROM brands LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	var brands []models.Brand

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to exectue statement: %w", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var brand models.Brand
		if err = rows.Scan(&brand.ID, &brand.Name); err != nil {
			return nil, fmt.Errorf("failed to execute rows: %w", err.Error())
		}
		brands = append(brands, brand)
	}
	return brands, nil
}
