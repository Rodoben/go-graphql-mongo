package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	uri := "mongodb://localhost:27017"
	fmt.Println(uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("unable to connect to mongo")
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil
	}
	fmt.Println("connected successfully")
	return &DB{
		client: client,
	}

}
