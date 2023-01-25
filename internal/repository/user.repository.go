package repository

import (
	"context"

	"github.com/romankravchuk/toronto-pizza/internal/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	db   *mongo.Database
	coll string
}

func NewUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		db:   db,
		coll: "users",
	}
}

func (r *MongoUserRepository) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	var (
		user = &models.User{}
		res  = r.db.Collection(r.coll).FindOne(ctx, bson.M{"phone": phone})
		err  = res.Decode(&user)
	)
	return user, err
}
