package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/seheraksam/shopshade_backend/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/user/register", handlers.CreateProductHandler)
	}
}
