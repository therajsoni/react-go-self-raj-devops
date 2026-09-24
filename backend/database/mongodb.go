package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

func ConnectMongoDB() {
	mongoURI := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(mongoURI),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB connected")

	UserCollection = client.Database("test").Collection("users_go")
}