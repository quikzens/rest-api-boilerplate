package db

import (
	"context"
	"log"

	"github.com/quikzens/rest-api-boilerplate/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DB Instance
var DB = func() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DbSource))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connecting to database")
	return client.Database(config.DbName)
}()

// Mongo Collection
var (
	UserColl = DB.Collection("users")
)
