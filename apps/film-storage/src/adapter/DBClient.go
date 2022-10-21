package adapter

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	conn *mongo.Client
}

type DBClient struct {
	mongo Mongo
	db    *mongo.Database
}

func MongoNewConnection() (db *Mongo) {
	host := os.Getenv("MONGODB_HOST")
	options := options.
		Client().
		ApplyURI(host)
	if os.Getenv("MONGODB_MONITOR") == "true" {
		cmdMonitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				log.Print(evt.Command)
			},
		}
		options.SetMonitor(cmdMonitor)
	}
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatalf("Error connectiong to mongo database: %s", err)
	}
	return &Mongo{conn: client}
}

func DBNewConnection() (dbClient *DBClient) {
	mongo := MongoNewConnection()

	mongoDBName := os.Getenv("MONGODB_DB_NAME")
	db := mongo.conn.Database(mongoDBName)
	return &DBClient{db: db, mongo: *mongo}
}

func (client *DBClient) Collection(name string) *mongo.Collection {
	return client.db.Collection(name)
}

func (client *DBClient) Ping() error {
	return client.mongo.conn.Ping(context.Background(), readpref.Primary())
}

func (client *DBClient) Close() {
	client.mongo.conn.Disconnect(context.Background())
	log.Println("closing db connection")
}
