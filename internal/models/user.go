package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	Username    string             `bson:"username"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Address     []Address          `bson:"address"`
	Phone       string             `bson:"phone"`
	CreatedAt   time.Time          `bson:"created_at"`
	Role        string             `bson:"role"`
	AuthToken   string             `bson:"auth_token"`
	ProviderId  primitive.ObjectID `bson:"provider_id"`
	UserType    UserType           `bson:"user_type"`
	TaxID       string             `bson:"tax_id"`
	CompanyName string             `bson:"company_name"`
}

type UserType string

const (
	Customer UserType = "customer"
	Admin    UserType = "admin"
	Seller   UserType = "seller"
)
