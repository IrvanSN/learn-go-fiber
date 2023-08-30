package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client

var mongoOnce sync.Once

var clientInstanceError error

type Collection string

const (
	ProductsCollection Collection = "products"
)

const (
	url      = "mongodb+srv://<user>:<password>@<host>/"
	Database = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOption := options.Client().ApplyURI(url)

		client, err := mongo.Connect(context.TODO(), clientOption)

		clientInstance = client

		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
