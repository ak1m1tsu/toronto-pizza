package repository

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepository struct {
	db   *mongo.Database
	coll string
}

func NewProductRepository(db *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		db:   db,
		coll: "products",
	}
}

func (r *MongoProductRepository) GetByID(ctx context.Context, id string) (*models.Product, error) {
	return nil, nil
}

func (r *MongoProductRepository) GetAll(ctx context.Context) ([]*models.Product, error) {
	return nil, nil
}

func (r *MongoProductRepository) Insert(ctx context.Context, user *models.Product) (*models.Product, error) {
	return nil, nil
}

func (r *MongoProductRepository) Update(ctx context.Context, id string, user *models.Product) (*models.Product, error) {
	return nil, nil
}

func (r *MongoProductRepository) Delete(ctx context.Context, id string) bool {
	return false
}
