package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/seheraksam/shopshade_backend/config"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMe(userID string) (models.User, error) {
	collection := config.GetCollection("shopshade", "users")
	fmt.Println("userID", userID)
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.User{}, errors.New("invalid user id")
	}
	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}
