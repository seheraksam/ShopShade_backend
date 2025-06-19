package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/seheraksam/shopshade_backend/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/product/create", handlers.CreateProductHandler)
		api.GET("/product/list", handlers.GetAllProductsHandler)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterHandler)
		auth.POST("/login", handlers.LoginHandler)
	}
}
