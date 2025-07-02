package services

import (
	"context"
	"fmt"
	"time"

	"github.com/seheraksam/shopshade_backend/config"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateWishList(ProductID string, UserID string) (string, error) {
	collection := config.GetCollection("shopshade", "wishlist")

	hex_productId, err := primitive.ObjectIDFromHex(ProductID)
	if err != nil {
		return "Product cannot", err
	}
	hex_userId, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		return "User cannot", err
	}

	wishlistData := models.Wishlist{
		ID:        primitive.NewObjectID(),
		ProductID: hex_productId,
		UserID:    hex_userId,
		CreatedAt: time.Now(),
	}
	fmt.Println(wishlistData)

	_, err = collection.InsertOne(context.TODO(), wishlistData)
	if err != nil {
		return "", err
	}

	return "Wishlist saved success", err
}

func DeleteWishList(WishlistID string) (string, error) {
	collection := config.GetCollection("shopshade", "wishlist")
	hex_wishlistsId, err := primitive.ObjectIDFromHex(WishlistID)
	if err != nil {
		return "User cannot", err
	}
	collection.DeleteOne(context.TODO(), bson.M{"_id": hex_wishlistsId})
	return "", nil
}
