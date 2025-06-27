package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodDebitCard    PaymentMethod = "debit_card"
	PaymentMethodCash         PaymentMethod = "cash"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)

type Payment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	OrderID       primitive.ObjectID `bson:"order_id"`
	Amount        float64            `bson:"amount"`
	PaymentMethod PaymentMethod      `bson:"payment_method"`
	Status        Status             `bson:"status"`
	CreatedAt     time.Time          `bson:"created_at"`
}
