package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProducts = "products"
)

type Products struct {
	ID          primitive.ObjectID `bson:"_id" json:"-"`
	Title       string             `bson:"title" form:"title" binding:"required" json:"title"`
	Description string             `bson:"description" json:"description"`
	Price       int                `bson:"price" json:"price"`
	SalePrice   int                `bson:"saleprice" json:"saleprice"`
	Status      string             `bson:"status" json:"status"`
	UserID      primitive.ObjectID `bson:"userID" json:"-"`
}

type ProductsRepository interface {
	Create(c context.Context, task *Products) error
	FetchByUserID(c context.Context, userID string) ([]Products, error)
}

type ProductsUsecase interface {
	Create(c context.Context, task *Products) error
	FetchByUserID(c context.Context, userID string) ([]Products, error)
}
