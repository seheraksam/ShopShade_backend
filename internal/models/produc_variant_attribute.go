package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductVariantAttribute struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ProductVariantID primitive.ObjectID `bson:"product_variant_id"`
	AttributeID      primitive.ObjectID `bson:"attribute_id"`
	Value            string             `bson:"value"`
}
