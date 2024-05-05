package storage

import (
	"context"
	"fmt"

	"github.com/tomef96/coop/mastodon/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s", config.MONGO_USER, config.MONGO_PASS, config.MONGO_URL)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	return client
}
