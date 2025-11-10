package handlers

import (
	"net/http"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) categories(c *gin.Context) {
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
