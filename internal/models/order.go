package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id" `
	UserID            primitive.ObjectID `bson:"user_id" json:"user_id" binding:"required"`
	OrderDate         time.Time          `bson:"order_date" json:"order_date" binding:"required"`
	TotalAmount       float64            `bson:"total_amount" json:"total_amount" binding:"required"`
	OrderStatus       primitive.ObjectID `bson:"order_status" json:"order_status" binding:"required"`
	PaymentID         primitive.ObjectID `bson:"payment_id" json:"payment_id" binding:"required"`
	ShippingAddressID primitive.ObjectID `bson:"shipping_address_id" json:"shipping_address_id" binding:"required"`
	BillAddressID     primitive.ObjectID `bson:"bill_address_id" json:"bill_address_id" binding:"required"`
	Items             []OrderItems       `bson:"order_items" json:"order_items" binding:"required"`
}
