package services

import (
	"github.com/seheraksam/shopshade_backend/internal/models"
	"github.com/seheraksam/shopshade_backend/internal/repositories"
)

func CreateNewProduct(product models.Product) error {
	_, err := repositories.CreateProduct(product)
	return err
}

func ListProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}
