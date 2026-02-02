package http_server

import (
	"e_shop_backend.amirkharisov.net/internal/http_server/middlewares"
	"e_shop_backend.amirkharisov.net/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

type authHandler interface {
	Login(c *gin.Context)
}

type categoryHandler interface {
	CreateCategory(c *gin.Context)
	Category(ctx *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetAllCategory(c *gin.Context)
}

type brandHandler interface {
	CreateBrand(c *gin.Context)
	Brand(ctx *gin.Context)
	UpdateBrand(c *gin.Context)
	DeleteBrand(c *gin.Context)
	GetAllBrand(c *gin.Context)
}

type productHandler interface {
	CreateProduct(c *gin.Context)
	Product(ctx *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetAllProduct(c *gin.Context)
}

type handlerAbs interface {
	authHandler
	categoryHandler
	brandHandler
	productHandler
}

func NewRouter(cfg *config.Config, handler handlerAbs) *gin.Engine {
	// Initialize default gin
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	authenticated := router.Group("/")
	authenticated.Use(middlewares.CheckAuth())

	// categories
	cg := authenticated.Group("/category")
	cg.POST("", handler.CreateCategory)
	cg.GET(":id", handler.Category)
	cg.PUT("", handler.UpdateCategory)
	cg.DELETE(":id", handler.DeleteCategory)
	cg.GET("all", handler.GetAllCategory)

	// brands
	bg := authenticated.Group("/brand")
	bg.POST("", handler.CreateBrand)
	bg.GET(":id", handler.Brand)
	bg.PUT("", handler.UpdateBrand)
	bg.DELETE(":id", handler.DeleteBrand)
	bg.GET("all", handler.GetAllBrand)

	// products
	pg := authenticated.Group("/product")
	pg.POST("", handler.CreateProduct)
	pg.GET(":id", handler.Product)
	pg.PUT("", handler.UpdateProduct)
	pg.DELETE(":id", handler.DeleteProduct)
	pg.GET("all", handler.GetAllProduct)

	public := router.Group("/")
	public.POST("/login", handler.Login)

	return router
}
