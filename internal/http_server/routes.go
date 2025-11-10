package http_server

import (
	"e_shop_backend.amirkharisov.net/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

type productHandler interface {
	Product(ctx *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetAllProduct(c *gin.Context)
}

type handlerAbs interface {
	productHandler
}

func NewRouter(cfg *config.Config, handler handlerAbs) *gin.Engine {
	// Initialize default gin
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Adding a category:
	// // * GET /categories
	// router.GET("/categories", handler.category)

	// * GET /product
	router.GET("/product/:id", handler.Product)

	// * POST /product
	router.POST("/product", handler.CreateProduct)

	router.PUT("/product", handler.UpdateProduct)

	router.DELETE("/product/:id", handler.DeleteProduct)

	router.GET("/product/all", handler.GetAllProduct)

	return router
}
