package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Order_Id  primitive.ObjectID `bson:"order_id,omitempty"`
	Type      Status             `bson:"type"`
	CreatedAt time.Time          `bson:"created_at"`
}

type Status string

const (
	OrderStatusPending    Status = "pending"
	OrderStatusProcessing Status = "processing"
	OrderStatusShipped    Status = "shipped"
	OrderStatusDelivered  Status = "delivered"
	OrderStatusCancelled  Status = "cancelled"
	OrderStatusRefunded   Status = "refunded"
	OrderStatusFailed     Status = "failed"
)

func IsValidOrderStatus(status string) bool {
	switch Status(status) {
	case OrderStatusPending, OrderStatusProcessing, OrderStatusShipped,
		OrderStatusDelivered, OrderStatusCancelled, OrderStatusRefunded, OrderStatusFailed:
		return true
	default:
		return false
	}
}
