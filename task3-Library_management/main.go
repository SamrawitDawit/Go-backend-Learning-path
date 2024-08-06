package main

import (
	"task3-Library_management/controllers"
	"task3-Library_management/models"
	"task3-Library_management/services"
)

func main() {

	libraryService := services.Library{}
	libraryController := controllers.LibraryController{
		LibraryService: &libraryService,
	}
	//sample
	libraryController.AddBook(models.Book{
		ID:     1,
		Title:  "The Courage to be Disliked",
		Author: "Ichiro Kishimi",
		Status: "Available",
	})
}
