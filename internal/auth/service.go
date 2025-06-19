package auth

import (
	"context"
	"errors"
	"time"

	"github.com/seheraksam/shopshade_backend/config"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user models.User) error {
	collection := config.GetCollection("shopshade", "users")

	var existing models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existing)
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()
	user.Role = "customer"
	user.UserType = models.Customer
	user.ProviderId = primitive.NewObjectID()

	_, err = collection.InsertOne(context.TODO(), user)
	return err
}

func LoginUser(email, password string) (models.User, error) {
	collection := config.GetCollection("shopshade", "users")

	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}
