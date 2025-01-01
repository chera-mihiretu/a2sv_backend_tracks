package main

import (
	"context"
	"fmt"
	"github/chera/task_manager/router"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() (mongo.Collection, error) {
	// Connect to database
	var collection mongo.Collection
	clientOoption := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(context.Background(), clientOoption)

	if err != nil {

		return collection, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Ping(ctx, nil)

	if err != nil {

		return collection, err
	}

	// Connect to collection
	collection = *client.Database(os.Getenv("DATABSE_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	fmt.Println("Database connected")
	return collection, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	collections, err := connectDB()
	if err != nil {
		log.Fatalf("error connecting to database %v", err.Error())
	}
	router.Routers(&collections)
}
