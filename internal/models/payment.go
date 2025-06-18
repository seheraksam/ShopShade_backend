package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	OrderID       primitive.ObjectID `bson:"order_id"`
	Amount        float64            `bson:"amount"`
	PaymentMethod string             `bson:"payment_method"`
	Status        string             `bson:"status"`
	CreatedAt     time.Time          `bson:"created_at"`
}
