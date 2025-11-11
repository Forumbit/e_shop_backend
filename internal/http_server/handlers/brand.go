package handlers

import (
	"net/http"
	"strconv"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBrand(c *gin.Context) {
	var brand models.Brand
	if err := c.ShouldBindBodyWithJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.brnd.Create(c, &brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
func (h *Handler) Brand(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	brand, err := h.brnd.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"brand": brand})
}
func (h *Handler) UpdateBrand(c *gin.Context) {
	var brand *models.Brand
	if err := c.ShouldBindBodyWithJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.brnd.Update(c, brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) DeleteBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = h.brnd.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) GetAllBrand(c *gin.Context) {
	brands, err := h.brnd.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"brands": brands})
}
