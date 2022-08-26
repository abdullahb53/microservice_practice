package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo db connection
func GetMongoDBCli() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://citizix:S3cret@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("mongodb_abdullah")
	collection = database.Collection("items")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	if collection == nil {
		err := database.CreateCollection(ctx, "items")
		if err != nil {
			log.Fatalf("Collection create failed:%v", err)
		}
	}

	//
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)

	/*
		List databases
	*/

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

}
