package main

import (
	"context"
	"log"
	"task5-Task_Mangement_API_using_MongoDB_and_Go_driver/controllers"
	service "task5-Task_Mangement_API_using_MongoDB_and_Go_driver/data"
	"task5-Task_Mangement_API_using_MongoDB_and_Go_driver/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	taskService := service.NewTaskService(client)
	taskController := &controllers.Controller{TaskService: *taskService}
	r := router.Router{Controller: taskController}
	r.Route()
}
