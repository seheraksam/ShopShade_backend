package services

import (
	"context"
	"fmt"

	"github.com/seheraksam/shopshade_backend/config"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CartItemAdd(cartItem models.CartItem) error {
	collection := config.GetCollection("shopshade", "cart_items")
	_, err := collection.InsertOne(context.TODO(), cartItem)
	if err != nil {
		return err
	}
	return nil
}

func CartItemUpdate(cartItem models.CartItem) error {
	collection := config.GetCollection("shopshade", "cart_items")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": cartItem.ID}, bson.M{"$set": cartItem})
	if err != nil {
		return err
	}
	return nil
}

func CartItemDelete(cartItem models.CartItem) error {
	collection := config.GetCollection("shopshade", "cart_items")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": cartItem.ID}, bson.M{"$set": bson.M{"is_deleted": true}})
	if err != nil {
		return err
	}
	fmt.Println("cartItem", cartItem)
	return nil
}

func GetUserCartItems(userID string) ([]models.CartItem, error) {
	collection := config.GetCollection("shopshade", "cart_items")
	fmt.Println("userID", userID)
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userObjectID, "is_deleted": false})
	if err != nil {
		return nil, err
	}
	fmt.Println("cursor", cursor)
	var cartItems []models.CartItem
	for cursor.Next(context.TODO()) {
		var cartItem models.CartItem
		if err := cursor.Decode(&cartItem); err != nil {
			return nil, err
		}
		fmt.Println("cartItem", cartItem)
		cartItems = append(cartItems, cartItem)
		fmt.Println("cartItems", cartItems)
	}
	return cartItems, nil
}

func GetDeletedUserCartItems(userID string) ([]models.CartItem, error) {
	collection := config.GetCollection("shopshade", "cart_items")
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userObjectID, "is_deleted": true})
	if err != nil {
		return nil, err
	}
	var cartItems []models.CartItem
	for cursor.Next(context.TODO()) {
		var cartItem models.CartItem
		if err := cursor.Decode(&cartItem); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}
