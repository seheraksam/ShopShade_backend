package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AttributeOptions struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	AttributeID primitive.ObjectID `bson:"attribute_id"`
	Value       string             `bson:"value"`
}
