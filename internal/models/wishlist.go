package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wishlist struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	Deleted   bool               `bson:"deleted" json:"deleted"`
}
