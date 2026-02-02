package handlers

import (
	"net/http"

	"e_shop_backend.amirkharisov.net/internal/domain/models"
	"e_shop_backend.amirkharisov.net/internal/pkg/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var u models.User
	if err := c.ShouldBindBodyWithJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if u.Username == "Check" && u.Password == "123456" {
		tokenString, err := auth.CreateToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "No username found")
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	}
}
