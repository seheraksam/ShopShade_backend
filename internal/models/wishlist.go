package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wishlist struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	UserID    primitive.ObjectID   `bson:"user_id"`
	ProductID []primitive.ObjectID `bson:"product_id"`
	CreatedAt time.Time            `bson:"created_at"`
}
