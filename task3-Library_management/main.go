package main

import (
	"fmt"
	"task3-Library_management/controllers"
	"task3-Library_management/models"
	"task3-Library_management/services"
)

func main() {

	libraryService := services.Library{}
	libraryController := controllers.LibraryController{
		LibraryService: &libraryService,
	}
	//sample book data
	libraryController.AddBook(models.Book{
		ID:     1,
		Title:  "The Courage to be Disliked",
		Author: "Ichiro Kishimi",
		Status: "Available",
	})
	libraryController.AddBook(models.Book{
		ID:     2,
		Title:  "Ego is the Enemy",
		Author: "Ryan Holiday",
		Status: "Available",
	})
	//testing ListAvailableBooks
	allBooks := libraryController.ListAvailableBooks()
	for _, book := range allBooks {
		fmt.Printf("book's id: %d\n", book.ID)
		fmt.Printf("book's name: %s\n", book.Title)
		fmt.Printf("book's author: %s\n", book.Author)
		fmt.Printf("book's status: %s\n", book.Status)
	}
	// sample member data
	libraryService.Members = make(map[int]*models.Member)
	libraryService.Members[1] = &models.Member{
		ID:            1,
		Name:          "Samrawit Dawit",
		BorrowedBooks: []models.Book{},
	}
	//testing borrow book
	libraryController.BorrowBook(2, 1)
	//testing list borrowed books
	borrowedBySamri := libraryController.ListBorrowedBooks(1)
	fmt.Printf("Borrowed by samri: %v\n", borrowedBySamri)
	//testing return book
	libraryController.ReturnBook(2, 1)
	borrowedBySamri = libraryController.ListBorrowedBooks(1)
	fmt.Printf("Borrowed by samri: %v", borrowedBySamri)
	//testing remove book
	libraryController.RemoveBook(1)
	allBooks = libraryController.ListAvailableBooks()
	for _, book := range allBooks {
		fmt.Printf("book's id: %d\n", book.ID)
		fmt.Printf("book's name: %s\n", book.Title)
		fmt.Printf("book's author: %s\n", book.Author)
		fmt.Printf("book's status: %s\n", book.Status)
	}
}
