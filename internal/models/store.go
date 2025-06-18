package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID          primitive.ObjectID `json:"id"`
	UserID      primitive.ObjectID `json:"user_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	LogoURL     string             `json:"logo_url"`
	BannerURL   string             `json:"banner_url"`
	IsVerified  bool               `json:"is_verified"`
	CreatedAt   time.Time          `json:"created_at"`
}
