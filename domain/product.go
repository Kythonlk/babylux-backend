package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProduct = "product"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"-"`
	Title       string             `bson:"title" form:"title" binding:"required" json:"title"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	SalePrice   float64            `bson:"salePrice" json:"salePrice"`
	Stock       int                `bson:"stock" json:"stock"`
	Review      int                `bson:"review" json:"review"`
	UserID      primitive.ObjectID `bson:"userID" json:"-"`
}

type ProductRepository interface {
	Create(c context.Context, product *Product) error
	FetchByUserID(c context.Context, userID string) ([]Product, error)
}

type ProductUsecase interface {
	Create(c context.Context, product *Product) error
	FetchByUserID(c context.Context, userID string) ([]Product, error)
}
