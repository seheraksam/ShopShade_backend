package repositories

import (
	"context"

	"github.com/seheraksam/shopshade_backend/config"
	"github.com/seheraksam/shopshade_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressCollection *mongo.Collection = config.GetCollection("shopshade", "address")

func CreateAddress(address models.Address) (*mongo.InsertOneResult, error) {
	return addressCollection.InsertOne(context.Background(), address)
}

func UpdateAddress(address models.Address) (*mongo.UpdateResult, error) {
	return addressCollection.UpdateOne(context.Background(), bson.M{"_id": address.ID}, bson.M{"$set": address})
}

func DeleteAddress(addressId string) (*mongo.DeleteResult, error) {
	return addressCollection.DeleteOne(context.Background(), bson.M{"_id": addressId})
}

func GetAddressById(addressId string) (*models.Address, error) {
	var address models.Address
	err := addressCollection.FindOne(context.Background(), bson.M{"_id": addressId}).Decode(&address)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func GetAddressByUserId(userId string) ([]models.Address, error) {
	cursor, err := addressCollection.Find(context.Background(), bson.M{"user_id": userId, "is_default": true})
	if err != nil {
		return nil, err
	}

	var addresses []models.Address
	if err = cursor.All(context.Background(), &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}

func GetDefaultAddressByUserId(userId string) (models.Address, error) {
	cursor, err := addressCollection.Find(context.Background(), bson.M{"user_id": userId, "is_default": true})
	if err != nil {
		return models.Address{}, err
	}

	var address models.Address
	if err = cursor.All(context.Background(), &address); err != nil {
		return models.Address{}, err
	}

	return address, nil
}
