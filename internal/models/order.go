package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	UserID            primitive.ObjectID `bson:"user_id"`
	OrderDate         time.Time          `bson:"order_date"`
	TotalAmount       float64            `bson:"total_amount"`
	OrderStatus       primitive.ObjectID `bson:"order_status"`
	PaymentID         primitive.ObjectID `bson:"payment_id"`
	ShippingAddressID primitive.ObjectID `bson:"shipping_address_id"`
	BillAddressID     primitive.ObjectID `bson:"bill_address_id"`
}
