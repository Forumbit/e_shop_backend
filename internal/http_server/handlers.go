package httpserver

import (
	"net/http"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func categories(c *gin.Context) {
	categories := []*models.Category{
		&models.Category{ID: 1, Name: "Electronics"},
		&models.Category{ID: 2, Name: "Clothing"},
		&models.Category{ID: 3, Name: "Books"},
	}
	var categoryNames []string
	for _, category := range categories {
		categoryNames = append(categoryNames, category.Name)
	}
	c.JSON(http.StatusOK, gin.H{"categories": categoryNames})
}

func product(c *gin.Context) {
	product := &models.Product{
		ID:                 1,
		Title:              "Sample Product",
		Description:        "This is a sample product.",
		Rating:             4.5,
		Category:           &models.Category{ID: 1, Name: "Electronics"},
		Brand:              &models.Brand{ID: 1, Name: "Brand A"},
		Price:              100,
		DiscountPercentage: 10,
		Tags:               []string{"tag1", "tag2"},
		Weight:             500,
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}
