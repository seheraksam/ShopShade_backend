package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderItems struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	OrderID         primitive.ObjectID `bson:"order_id"`
	ProductID       primitive.ObjectID `bson:"product_id"`
	Quantity        int                `bson:"quantity"`
	PriceAtPurchase float64            `bson:"price_at_purchase"`
	StoreID         primitive.ObjectID `bson:"store_id"`
}
