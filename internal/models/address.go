package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	AddressType AddressType        `bson:"address_type"`
	Title       string             `bson:"title"`
	Line1       string             `bson:"line1"`
	Line2       string             `bson:"line2"`
	City        string             `bson:"city"`
	State       string             `bson:"state"`
	ZipCode     string             `bson:"zip_code"`
	Country     string             `bson:"country"`
	IsDefault   bool               `bson:"is_default"`
	IsActive    bool               `bson:"is_active"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type AddressType string

const (
	Invoice  AddressType = "invoice"
	Delivery AddressType = "delivery"
)
