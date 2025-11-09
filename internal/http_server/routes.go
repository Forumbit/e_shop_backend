package httpserver

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	// Initialize default gin
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Adding a category:
	// * GET /categories
	router.GET("/categories", categories)

	// * GET /product
	router.GET("/product", product)

	return router
}
