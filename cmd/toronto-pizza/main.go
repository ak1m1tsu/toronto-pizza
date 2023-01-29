package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/romankravchuk/toronto-pizza/internal/config"
	"github.com/romankravchuk/toronto-pizza/internal/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	conf       = &config.Config{}
	client     = &mongo.Client{}
	db         = &mongo.Database{}
	err        error
	configPath string
)

func init() {
	configPath = os.Getenv("CONFIG_FILE_PATH")
	conf, err = config.GetConfig(configPath)
	if err != nil {
		panic(err.Error())
	}

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.MongoURL))
	if err != nil {
		panic(err.Error())
	}

	db = client.Database("toronto-pizza")
}

func main() {
	r := router.NewRouter(db, conf)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), r))
}
