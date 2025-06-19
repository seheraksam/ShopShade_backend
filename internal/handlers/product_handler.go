package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"github.com/seheraksam/shopshade_backend/internal/services"
)

func CreateProductHandler(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateNewProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ürün eklenemedi"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ürün başarıyla eklendi"})
}

func GetAllProductsHandler(c *gin.Context) {
	products, err := services.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ürünler alınamadı"})
		return
	}

	c.JSON(http.StatusOK, products)
}
