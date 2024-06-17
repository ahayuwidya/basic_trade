package routers

import (
	"basic_trade/controllers"
	"basic_trade/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.AdminRegister)
		authRouter.POST("/login", controllers.AdminLogin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", controllers.GetProduct)
		productRouter.GET("/:productUUID", controllers.GetProductbyUUID)

		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productUUID", middlewares.ProductAuthorization(), controllers.UpdateProductbyUUID)
		productRouter.DELETE((":productUUID"), middlewares.ProductAuthorization(), controllers.DeleteProductbyUUID)
	}

	variantRouter := router.Group("/variants")
	{
		variantRouter.GET("/", controllers.GetVariant)
		variantRouter.GET("/:variantUUID", controllers.GetVariantbyUUID)

		variantRouter.Use(middlewares.Authentication())
		variantRouter.POST("/", controllers.CreateVariant)
		variantRouter.PUT("/:variantUUID", middlewares.VariantAuthorization(), controllers.UpdateVariantbyUUID)
		variantRouter.DELETE("/:variantUUID", middlewares.VariantAuthorization(), controllers.DeleteVariantbyUUID)
	}

	return router
}
