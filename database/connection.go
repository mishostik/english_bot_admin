package database

import (
	"context"
	"log"

	//"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	dbase  *mongo.Database
}

func NewConnection(ctx context.Context, dbName string, URI string) (*Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal("Error while connection to MongoDB")
	}
	db := client.Database(dbName)
	return &Database{
		client: client,
		dbase:  db,
	}, nil
}

func (db *Database) Collection(collectionName string) (*mongo.Collection, error) {
	return db.dbase.Collection(collectionName), nil
}

func (db *Database) Disconnect() error {
	return db.client.Disconnect(context.Background())
}
