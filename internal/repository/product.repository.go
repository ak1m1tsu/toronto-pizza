package repository

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository/filter"
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

func (r *MongoProductRepository) GetAll(ctx context.Context, filter *filter.ProductFilter, sort *filter.ProductSort, page int) ([]*models.Product, error) {
	aggregatePipline := r.buildProductPipline(filter, sort, page)
	cursor, err := r.db.Collection(r.coll).Aggregate(ctx, aggregatePipline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var products []*models.Product
	err = cursor.All(ctx, &products)
	return products, err
}

func (r *MongoProductRepository) buildProductPipline(filter *filter.ProductFilter, sort *filter.ProductSort, page int) []bson.M {
	aggregatePipline := make([]bson.M, 0)
	if filter != nil {
		filterDoc := r.buildProductFilter(filter)
		if filterDoc != nil && len(filterDoc) > 0 {
			aggregatePipline = append(aggregatePipline, bson.M{"$match": filterDoc})
		}
	}
	if sort != nil {
		sortDoc := r.buildProductSort(sort)
		if sortDoc != nil && len(sortDoc) > 0 {
			aggregatePipline = append(aggregatePipline, bson.M{"$sort": sortDoc})
		}
	}
	aggregatePipline = append(aggregatePipline, bson.M{"$skip": int(10 * (page - 1))})
	aggregatePipline = append(aggregatePipline, bson.M{"$limit": 10})
	return aggregatePipline
}

func (r *MongoProductRepository) buildProductFilter(filter *filter.ProductFilter) bson.M {
	filterDoc := bson.M{}
	if filter.Category != "" {
		filterDoc["category"] = filter.Category
	}
	if filter.PriceMin != 0 || filter.PriceMax != 0 {
		priceFilter := bson.M{}
		if filter.PriceMin != 0 {
			priceFilter["$gte"] = filter.PriceMin
		}
		if filter.PriceMax != 0 {
			priceFilter["$lte"] = filter.PriceMax
		}
		filterDoc["price"] = priceFilter
	}
	if filter.Name != "" {
		filterDoc["name"] = bson.M{"$regex": primitive.Regex{Pattern: filter.Name, Options: "i"}}
	}
	return filterDoc
}

func (r *MongoProductRepository) buildProductSort(sort *filter.ProductSort) bson.M {
	sortDoc := bson.M{}
	for _, opt := range sort.Options {
		if opt.Field != "" {
			sortOrder := 1
			if opt.Order == filter.Descending {
				sortOrder = -1
			}
			sortDoc[opt.Field] = sortOrder
		}
	}
	return sortDoc
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
