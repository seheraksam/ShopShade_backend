package repositories

import (
	"context"

	"github.com/seheraksam/shopshade_backend/config"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = config.GetCollection("shopshade", "product")

func CreateProduct(product models.Product) (*mongo.InsertOneResult, error) {
	return productCollection.InsertOne(context.Background(), product)
}

func GetAllProducts() ([]models.Product, error) {
	cursor, err := productCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var products []models.Product
	if err = cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}

	return products, nil
}
