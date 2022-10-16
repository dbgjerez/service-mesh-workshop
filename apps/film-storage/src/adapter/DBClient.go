package adapter

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	uri = "mongodb://localhost:27017"
)

type DBClient struct {
	db *mongo.Client
}

func DBNewConnection() (dbClient *DBClient) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("error connectiong to mongo database: %s", err)
	}
	return &DBClient{db: client}
}

func (client *DBClient) Ping() error {
	return client.db.Ping(context.TODO(), readpref.Primary())
}

func (client *DBClient) Close() {
	client.db.Disconnect(context.TODO())
	log.Println("closing db connection")
}
