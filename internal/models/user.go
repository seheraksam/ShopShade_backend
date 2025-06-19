package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string             `bson:"firstname" json:"first_name"`
	LastName    string             `bson:"lastname" json:"last_name"`
	Username    string             `bson:"username" json:"username"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password"`
	Address     []Address          `bson:"address" json:"address"`
	Phone       string             `bson:"phone" json:"phone"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	Role        string             `bson:"role" json:"role"`
	AuthToken   string             `bson:"auth_token" json:"auth_token"`
	ProviderId  primitive.ObjectID `bson:"provider_id" json:"provider_id"`
	UserType    UserType           `bson:"user_type" json:"user_type"`
	TaxID       string             `bson:"tax_id" json:"tax_id"`
	CompanyName string             `bson:"company_name" json:"company_name"`
}

type UserType string

const (
	Customer UserType = "customer"
	Admin    UserType = "admin"
	Seller   UserType = "seller"
)
