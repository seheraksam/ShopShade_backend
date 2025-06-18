package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductAttribute struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProductID   primitive.ObjectID `bson:"product_id"`
	AttributeID primitive.ObjectID `bson:"attribute_id"`
	Value       string             `bson:"value"`
}
