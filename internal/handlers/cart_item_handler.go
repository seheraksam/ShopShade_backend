package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"github.com/seheraksam/shopshade_backend/internal/services"
)

func CartItemAddHandler(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartItem.CreatedAt = time.Now()
	cartItem.UpdatedAt = time.Now()
	cartItem.IsDeleted = false
	err := services.CartItemAdd(cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart item added successfully"})
}

func CartItemUpdateHandler(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CartItemUpdate(cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart item updated successfully"})
}

func CartItemDeleteHandler(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CartItemDelete(cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart item deleted successfully"})
}

func CartItemListHandler(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("req", req)
	userID := req.UserID
	fmt.Println("userID", userID)
	cartItems, err := services.GetUserCartItems(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartItems)
}
