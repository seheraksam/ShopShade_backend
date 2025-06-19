package services

import (
	"github.com/seheraksam/shopshade_backend/internal/models"
	"github.com/seheraksam/shopshade_backend/internal/repositories"
)

func CreateAddress(address models.Address) error {
	_, err := repositories.CreateAddress(address)
	return err
}

func UserAddress(address models.Address, userId string) ([]models.Address, error) {
	return repositories.GetAddressByUserId(userId)
}

func GetDefaultAddress(userId string) (models.Address, error) {
	return repositories.GetDefaultAddressByUserId(userId)
}
