package main

import (
	"context"
	"log"
	router "task8-Testing/Delivery/routers"
	infrastructure "task8-Testing/Infrastructure"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	jwt_service := &infrastructure.JWT_Service{}

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("taskdb")
	router.NewTaskRouter(10*time.Second, *db, jwt_service)
}
