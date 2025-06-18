package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Attribute struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	CategoryID primitive.ObjectID `bson:"category_id"`
	IsVariant  bool               `bson:"is_variant"`
	InputType  string             `bson:"input_type"`
}
