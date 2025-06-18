package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	CategoryID  primitive.ObjectID `bson:"category_id"`
	StoreID     primitive.ObjectID `bson:"store_id"`
	Description string             `bson:"description"`
	Weight      float64            `bson:"weight,omitempty"`
	IsActive    bool               `bson:"is_active"`
	Gender      string             `bson:"gender"`
	Brand       string             `bson:"brand"`
}
