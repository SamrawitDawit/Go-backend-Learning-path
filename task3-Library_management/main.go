package main

import (
	"task3-Library_management/controllers"
	"task3-Library_management/models"
	"task3-Library_management/services"
)

func main() {
	//dummy data for members
	members := make(map[int]*models.Member)
	members[1] = &models.Member{ID: 1, Name: "Samri", BorrowedBooks: make([]models.Book, 0)}
	members[2] = &models.Member{ID: 2, Name: "Redi", BorrowedBooks: make([]models.Book, 0)}
	library := services.NewLibrary(make(map[int]*models.Book), members)
	libController := controllers.LibraryController{LibraryManager: library}
	libController.StartLibraryController()
}
