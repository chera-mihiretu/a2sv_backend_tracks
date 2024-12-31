package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	const url = "mongodb://localhost:27017"
	clientOption := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	trainer := Trainer{"Abdi Esayaas", 22, "Adama"}

	collection := client.Database("test").Collection("trainers")

	insertRestult, err := collection.InsertOne(context.TODO(), trainer)
	if err != nil {
		log.Fatalf("Failed to insert data %v", err)
	}

	fmt.Printf("Data is inserted with ID : %v \n", insertRestult.InsertedID)

	// here we display all the data

	result, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("Failed to display datas %v\n", err)
	}

	for result.Next(context.TODO()) {
		var trainer Trainer
		if err := result.Decode(&trainer); err != nil {
			log.Fatalf("Error Log %s \n", err)
		}

		fmt.Println(trainer)
	}
	var new_t Trainer
	result, err = collection.Find(context.TODO(), bson.D{{Key: "name", Value: "Abdi Esayaas"}})

	if err != nil {
		log.Fatalf("Error occured %s \n", err)
	}
	fmt.Println("=====================================")
	for result.Next(context.TODO()) {
		var t Trainer

		if err := result.Decode(&t); err != nil {
			log.Fatal("Error happened")
		}

		fmt.Println(t)
	}
	fmt.Println(new_t)
	client.Disconnect(context.TODO())

}
