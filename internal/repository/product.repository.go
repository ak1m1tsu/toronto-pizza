package repository

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var (
		product  = &models.Product{}
		objID, _ = primitive.ObjectIDFromHex(id)
		filter   = bson.M{"_id": objID}
		result   = r.db.Collection(r.coll).FindOne(ctx, filter)
		err      = result.Decode(product)
	)
	return product, err
}

func (r *MongoProductRepository) GetAll(ctx context.Context) ([]*models.Product, error) {
	cursor, err := r.db.Collection(r.coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}
	var products []*models.Product
	err = cursor.All(ctx, &products)
	return products, err
}

func (r *MongoProductRepository) Insert(ctx context.Context, product *models.Product) (*models.Product, error) {
	res, err := r.db.Collection(r.coll).InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	product.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return product, err
}

func (r *MongoProductRepository) Update(ctx context.Context, id string, p *models.Product) (*models.Product, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		filter   = bson.M{"_id": objID}
		update   = bson.M{"$set": bson.M{"name": p.Name, "description": p.Description, "price": p.Price, "category": p.Category}}
		_, err   = r.db.Collection(r.coll).UpdateOne(ctx, filter, update)
	)
	return p, err
}

func (r *MongoProductRepository) Delete(ctx context.Context, id string) error {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		filter   = bson.M{"_id": objID}
		_, err   = r.db.Collection(r.coll).DeleteOne(ctx, filter)
	)
	return err
}
