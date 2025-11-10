package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
	"github.com/lib/pq"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) Create(ctx context.Context, product *models.Product) (int, error) {
	stmt, err := p.db.Prepare("INSERT INTO products (title, description, rating, category_id, brand_id, price, discount, tags, weight) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRowContext(ctx, product.Title, product.Description, product.Rating, product.CategoryID, product.BrandID, product.Price, product.Discount, pq.Array(product.Tags), product.Weight).Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}

	return id, nil
}

func (p *ProductRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	stmt, err := p.db.Prepare("SELECT * FROM products WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var pr models.Product

	if err = stmt.QueryRow(id).Scan(&pr.ID, &pr.Title, &pr.Description, &pr.Rating, &pr.CategoryID, &pr.BrandID, &pr.Price, &pr.Discount, pq.Array(&pr.Tags), &pr.Weight); err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	return &pr, nil
}

func (p *ProductRepository) Update(ctx context.Context, pr *models.Product) error {
	stmt, err := p.db.Prepare("UPDATE products SET title = $2, description = $3, rating = $4, category_id = $5, brand_id = $6, price = $7, discount = $8, tags = $9, weight = $10 WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()
	if _, err = stmt.Exec(pr.ID, pr.Title, pr.Description, pr.Rating, pr.CategoryID, pr.BrandID, pr.Price, pr.Discount, pq.Array(pr.Tags), pr.Weight); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err.Error())
	}
	return nil
}

func (p *ProductRepository) Delete(ctx context.Context, id int) error {
	stmt, err := p.db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()
	if _, err = stmt.Exec(id); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err.Error())
	}
	return nil
}

func (p *ProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	stmt, err := p.db.Prepare("SELECT * FROM products LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err.Error())
	}
	defer stmt.Close()

	var rows *sql.Rows
	var products []models.Product

	if rows, err = stmt.Query(); err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var pr models.Product
		err = rows.Scan(&pr.ID, &pr.Title, &pr.Description, &pr.Rating, &pr.CategoryID, &pr.BrandID, &pr.Price, &pr.Discount, pq.Array(&pr.Tags), &pr.Weight)
		if err != nil {
			return nil, fmt.Errorf("failed to scan structs: %w", err)
		}
		products = append(products, pr)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over products: %w", err)
	}

	return products, nil
}
