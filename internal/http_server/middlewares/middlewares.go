package middlewares

import (
	"net/http"

	"e_shop_backend.amirkharisov.net/internal/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || len(tokenString) <= 7 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		err := auth.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
