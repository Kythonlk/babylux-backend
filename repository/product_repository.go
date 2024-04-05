package repository

import (
	"context"

	"go-backend/domain"
	"go-backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productRepository struct {
	database   mongo.Database
	collection string
}

func NewProductRepository(db mongo.Database, collection string) domain.ProductRepository {
	return &productRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *productRepository) Create(c context.Context, product *domain.Product) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, product)

	return err
}

func (tr *productRepository) FetchByUserID(c context.Context, userID string) ([]domain.Product, error) {
	collection := tr.database.Collection(tr.collection)

	var products []domain.Product

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return products, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &products)
	if products == nil {
		return []domain.Product{}, err
	}

	return products, err
}
