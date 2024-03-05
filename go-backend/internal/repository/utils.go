package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DB *mongo.Database

func Connect2MongoDB() (*mongo.Client, error) {
	defer log.Println("Connected to MongoDB!")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions) // use context.Background()
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil) // use context.Background()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Disconnect2MongoDB(client *mongo.Client, ctx context.Context) {
	defer log.Println("Connection to MongoDB closed.")

	err := client.Disconnect(ctx)
	if err != nil {
		log.Println(err)
	}
}

func SelectDB(client *mongo.Client, dbName string) *mongo.Database {
	return client.Database(dbName)
}

func SelectCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
