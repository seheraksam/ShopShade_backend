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
		api.POST("/login", handlers.LoginHandler)
		api.POST("/register", handlers.RegisterHandler)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterHandler)
		auth.POST("/login", handlers.LoginHandler)
	}
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
