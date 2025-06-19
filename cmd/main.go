package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seheraksam/shopshade_backend/internal/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
