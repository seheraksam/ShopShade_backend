package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	CategoryID  primitive.ObjectID `bson:"category_id" json:"category_id"`
	StoreID     primitive.ObjectID `bson:"store_id" json:"store_id"`
	Description string             `bson:"description" json:"description"`
	Weight      float64            `bson:"weight,omitempty" json:"weight"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	Gender      string             `bson:"gender" json:"gender"`
	Brand       string             `bson:"brand" json:"brand"`
}
