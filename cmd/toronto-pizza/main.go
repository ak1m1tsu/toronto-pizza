package main

import (
	"context"
	"net/http"
	"os"

	"github.com/romankravchuk/toronto-pizza/internal/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client = &mongo.Client{}
	db     = &mongo.Database{}
	err    error
)

func init() {
	config := LoadConfig()
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}
	db = client.Database("toronto-pizza")
}

func main() {
	r := router.NewRouter(db)
	http.ListenAndServe(":3000", r)
}

type config struct {
	MongoURI string
}

func LoadConfig() *config {
	return &config{
		MongoURI: os.Getenv("MONGO_URI"),
	}
}
