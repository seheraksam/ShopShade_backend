package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/shopshade_backend/config"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"github.com/seheraksam/shopshade_backend/internal/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrderHandler(c *gin.Context) {
	var req struct {
		UserID            string         `json:"user_id"`
		ShippingAddressID string         `json:"shipping_address_id"`
		BillAddressID     string         `json:"bill_address_id"`
		Payment           models.Payment `json:"payment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	selectedItems, err := services.GetUserCartItems(req.UserID)
	if err != nil || len(selectedItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "seçili ürün yok"})
		return
	}

	var total float64
	for _, item := range selectedItems {
		total += float64(item.Quantity) * item.Price
	}

	payment := models.Payment{
		ID:            req.Payment.ID,
		PaymentMethod: req.Payment.PaymentMethod,
		Status:        req.Payment.Status,
		Amount:        req.Payment.Amount,
		CreatedAt:     time.Now(),
	}

	config.GetCollection("shopshade", "payments").InsertOne(c, payment)

	status := models.OrderStatus{
		ID:        primitive.NewObjectID(),
		Type:      "pending",
		CreatedAt: time.Now(),
	}
	config.GetCollection("shopshade", "order_statuses").InsertOne(c, status)

	userID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	var orderItems []models.OrderItems
	for _, cartItem := range selectedItems {
		orderItems = append(orderItems, models.OrderItems{
			ID:        primitive.NewObjectID(),
			OrderID:   primitive.NewObjectID(),
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		})
	}
	hxBillAddressID, _ := primitive.ObjectIDFromHex(req.BillAddressID)
	order := models.Order{
		ID:                primitive.NewObjectID(),
		UserID:            userID,
		OrderDate:         time.Now(),
		TotalAmount:       total,
		PaymentID:         payment.ID,
		ShippingAddressID: primitive.NilObjectID,
		BillAddressID:     hxBillAddressID,
		Items:             orderItems,
	}
	if req.ShippingAddressID != "" {
		order.ShippingAddressID, _ = primitive.ObjectIDFromHex(req.ShippingAddressID)
	}
	if req.BillAddressID != "" {
		order.BillAddressID, _ = primitive.ObjectIDFromHex(req.BillAddressID)
	}

	config.GetCollection("shopshade", "orders").InsertOne(c, order)

	orderItemsColl := config.GetCollection("shopshade", "order_items")
	for _, item := range orderItems {
		orderItemsColl.InsertOne(c, item)
	}

	err = services.MarkSelectedCartItemsDeleted(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cart cleanup failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sipariş oluşturuldu", "order_id": order.ID.Hex()})
}

func GetOrderHandler(c *gin.Context) {
	var req struct {
		OrderID string `json:"order_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := services.GetOrder(req.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetUserOrdersHandler(c *gin.Context) {
	userIDHex := c.GetString("user_id")
	if userIDHex == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	orders, err := services.GetUserOrders(userIDHex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Siparişler alınamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func UpdateOrderStatusHandler(c *gin.Context) {
	orderID := c.Param("id")
	var req struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || !models.IsValidOrderStatus(req.Status) {
		c.JSON(400, gin.H{"error": "Geçersiz durum"})
		return
	}

	objID, _ := primitive.ObjectIDFromHex(orderID)
	coll := config.GetCollection("shopshade", "orders")
	_, err := coll.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{
		"$set": bson.M{"order_status": req.Status},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Durum güncellenemedi"})
		return
	}

	c.JSON(200, gin.H{"message": "Durum güncellendi"})
}
