package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductVariant struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID    primitive.ObjectID `bson:"product_id"`
	Size         string             `bson:"size"`
	Color        string             `bson:"color"`
	Price        float64            `bson:"price"`
	Stock        int                `bson:"stock"`
	IsActive     bool               `bson:"is_active"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	IsDeleted    bool               `bson:"is_deleted"`
	IsAvailable  bool               `bson:"is_available"`
	IsFeatured   bool               `bson:"is_featured"`
	IsNew        bool               `bson:"is_new"`
	IsBestSeller bool               `bson:"is_best_seller"`
	IsTopRated   bool               `bson:"is_top_rated"`
	IsTrending   bool               `bson:"is_trending"`
	IsSale       bool               `bson:"is_sale"`
	IsHot        bool               `bson:"is_hot"`
	SKU          string             `bson:"sku"`
	Images       []string           `bson:"images"`
}
