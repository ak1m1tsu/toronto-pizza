// repository contains the implementation of a MongoDB repository for products
package repository

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository/filter"
	"github.com/romankravchuk/toronto-pizza/internal/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoProductRepository is a struct that implements product repository using MongoDB as storage.
type MongoProductRepository struct {
	// db is a reference to a MongoDB database.
	db *mongo.Database

	// coll is a string representing the name of the products collection in the database.
	coll string
}

// NewProductRepository creates a new instance of MongoProductRepository and returns a pointer to it.
func NewProductRepository(db *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		db:   db,
		coll: "products",
	}
}

// GetByID returns a product with a given ID.
//
// If a product with the given ID is not found, the error will be set to mongo.ErrNoDocuments.
//
// id is a string representation of the ObjectID of the product.
// ctx is a context used to cancel the request.
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

// GetAll returns a slice of all products that match the filter criteria, sorted according to the sort options, and paged.
// If there are no products matching the filter criteria, an empty slice will be returned.
//
// filter is a pointer to a ProductFilter that specifies the filter criteria for the query.
// sort is a pointer to a ProductSort that specifies the sort options for the query.
// page is an integer representing the number of the page to return.
// ctx is a context used to cancel the request.
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

// buildProductPipline creates a pipeline for aggregating products based on the filter and sort criteria, and the desired page number.
// The pipeline consists of multiple stages, including filtering based on the provided criteria, sorting the results, skipping to the desired page, and limiting the results to 10 products.
// The pipeline is represented as a slice of bson.M, which is a type used to represent a BSON document.
// The function returns the created pipeline.
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

// buildProductFilter creates a filter document for aggregating products based on the provided filter criteria.
// The filter document is represented as a bson.M, which is a type used to represent a BSON document.
// The function returns the created filter document.
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

// buildProductSort creates a BSON document for sorting products in a mongodb collection.
// It takes in a sort option (sort.ProductSort) and returns a BSON document (bson.M) to be used in the mongodb aggregation pipeline.
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

// Insert inserts a new product into the mongodb collection.
// It takes in a context (context.Context) and a product instance (models.Product), inserts it into the collection,
// and returns the inserted product instance with its ID set or an error.
func (r *MongoProductRepository) Insert(ctx context.Context, product *models.Product) (*models.Product, error) {
	res, err := r.db.Collection(r.coll).InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	product.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return product, err
}

// Update updates a product in the mongodb collection.
// It takes in a context (context.Context), the product ID (string), and a product instance (models.Product).
// It updates the product in the collection and returns the updated product instance or an error.
func (r *MongoProductRepository) Update(ctx context.Context, id string, p *models.Product) (*models.Product, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		filter   = bson.M{"_id": objID}
		update   = bson.M{"$set": bson.M{"name": p.Name, "description": p.Description, "price": p.Price, "category": p.Category}}
		_, err   = r.db.Collection(r.coll).UpdateOne(ctx, filter, update)
	)
	return p, err
}

// Delete deletes a product from the mongodb collection.
// It takes in a context (context.Context) and a product ID (string). It deletes the product from the collection and returns an error if there was an issue.
func (r *MongoProductRepository) Delete(ctx context.Context, id string) error {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		filter   = bson.M{"_id": objID}
		_, err   = r.db.Collection(r.coll).DeleteOne(ctx, filter)
	)
	return err
}
