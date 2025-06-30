package services

import (
	"context"
	"time"

	"github.com/seheraksam/shopshade_backend/config"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(order models.Order) (string, error) {
	collection := config.GetCollection("shopshade", "orders")
	order.OrderDate = time.Now()
	orderStatusCollection := config.GetCollection("shopshade", "order_statuses")
	orderStatus, err := orderStatusCollection.InsertOne(context.TODO(), bson.M{"name": "pending"})
	if err != nil {
		return "Order status not found", err
	}
	order.OrderStatus = orderStatus.InsertedID.(primitive.ObjectID)
	paymentCollection := config.GetCollection("shopshade", "payments")
	payment, err := paymentCollection.InsertOne(context.TODO(), bson.M{"name": "pending"})
	if err != nil {
		return "Payment not found", err
	}
	order.PaymentID = payment.InsertedID.(primitive.ObjectID)
	shippingAddressCollection := config.GetCollection("shopshade", "shipping_addresses")
	shippingAddress, err := shippingAddressCollection.InsertOne(context.TODO(), bson.M{"name": "pending"})
	if err != nil {
		return "Shipping address not found", err
	}
	order.ShippingAddressID = shippingAddress.InsertedID.(primitive.ObjectID)
	order.BillAddressID = primitive.ObjectID{}
	order.Items = []models.OrderItems{}
	_, err = collection.InsertOne(context.TODO(), order)
	if err != nil {
		return "Order not created", err
	}
	return "Order created successfully", nil
}

func GetOrder(orderID string) (models.Order, error) {
	collection := config.GetCollection("shopshade", "orders")
	orderObjectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return models.Order{}, err
	}
	order := models.Order{}

	err = collection.FindOne(context.TODO(), bson.M{"_id": orderObjectID}).Decode(&order)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func GetUserOrders(userID string) ([]models.Order, error) {
	collection := config.GetCollection("shopshade", "orders")
	HexUserId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": HexUserId})
	if err != nil {
		return nil, err
	}
	var orders []models.Order
	for cursor.Next(context.TODO()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func MarkSelectedCartItemsDeleted(userID string) error {
	collection := config.GetCollection("shopshade", "cart_items")
	_, err := collection.UpdateMany(context.TODO(), bson.M{"user_id": userID, "deleted": false}, bson.M{"$set": bson.M{"deleted": true}})
	if err != nil {
		return err
	}
	return nil
}
