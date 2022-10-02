package initializers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToDB() {
	var err error
	// Initialize MongoDB Client
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_URI))

	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	log.Printf("Connected to database")
}
