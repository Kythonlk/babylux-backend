package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProduct = "product"
)

type Product struct {
	ID     primitive.ObjectID `bson:"_id" json:"-"`
	Title  string             `bson:"title" form:"title" binding:"required" json:"title"`
	UserID primitive.ObjectID `bson:"userID" json:"-"`
}

type ProductRepository interface {
	Create(c context.Context, product *Product) error
	FetchByUserID(c context.Context, userID string) ([]Product, error)
}

type ProductUsecase interface {
	Create(c context.Context, product *Product) error
	FetchByUserID(c context.Context, userID string) ([]Product, error)
}
