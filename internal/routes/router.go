package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/seheraksam/shopshade_backend/internal/auth"
	"github.com/seheraksam/shopshade_backend/internal/handlers"
	"github.com/seheraksam/shopshade_backend/middleware"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/product/create", middleware.AuthMiddleware(), middleware.RequireRole("seller"), handlers.CreateProductHandler)
		api.GET("/product/list", middleware.AuthMiddleware(), handlers.GetAllProductsHandler)
		api.POST("/login", handlers.LoginHandler)
		api.POST("/register", middleware.AuthMiddleware(), handlers.RegisterHandler)
		api.GET("/profile", middleware.AuthMiddleware(), func(c *gin.Context) {
			userID := c.MustGet("user_id").(string)
			c.JSON(200, gin.H{"message": "Hoş geldin!", "user_id": userID})
		})
	}
	api.POST("/refresh", func(c *gin.Context) {
		var body struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Refresh token eksik"})
			return
		}

		token, err := jwt.Parse(body.RefreshToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Refresh token geçersiz"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		role := claims["role"].(string)

		newAccessToken, newRefreshToken, err := auth.GenerateJWT(userID, role)
		if err != nil {
			c.JSON(500, gin.H{"error": "Token üretilemedi"})
			return
		}

		c.JSON(200, gin.H{
			"access_token":  newAccessToken,
			"refresh_token": newRefreshToken,
		})
	})

}

//api.POST("/address/create", handlers.CreateAddressHandler)
//api.GET("/address/list", handlers.ListAddressHandler)
/*api.GET("/address/default", handlers.GetDefaultAddressHandler)
api.GET("/address/list", handlers.ListAddressHandler)
api.POST("/address/create", handlers.CreateAddressHandler)
api.PUT("/address/update", handlers.UpdateAddressHandler)
api.DELETE("/address/delete", handlers.DeleteAddressHandler)
api.POST("/order/create", handlers.CreateOrderHandler)
api.GET("/order/list", handlers.ListOrdersHandler)
api.GET("/order/detail", handlers.GetOrderDetailHandler)
api.PUT("/order/update", handlers.UpdateOrderHandler)
api.DELETE("/order/delete", handlers.DeleteOrderHandler)
api.POST("/payment/create", handlers.CreatePaymentHandler)
api.GET("/payment/list", handlers.ListPaymentsHandler)
api.GET("/payment/detail", handlers.GetPaymentDetailHandler)
api.PUT("/payment/update", handlers.UpdatePaymentHandler)
api.DELETE("/payment/delete", handlers.DeletePaymentHandler)
api.POST("/cart/add", handlers.AddToCartHandler)
api.GET("/cart/list", handlers.ListCartHandler)
api.DELETE("/cart/delete", handlers.DeleteCartHandler)
api.POST("/cart/checkout", handlers.CheckoutHandler)
api.POST("/cart/add", handlers.AddToCartHandler)
api.GET("/cart/list", handlers.ListCartHandler)
api.DELETE("/cart/delete", handlers.DeleteCartHandler)
api.POST("/cart/checkout", handlers.CheckoutHandler)
api.POST("/cart/add", handlers.AddToCartHandler)
api.GET("/cart/list", handlers.ListCartHandler)
api.DELETE("/cart/delete", handlers.DeleteCartHandler)*/
