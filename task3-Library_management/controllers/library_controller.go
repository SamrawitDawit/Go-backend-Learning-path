package controllers

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"task3-Library_management/models"
	"task3-Library_management/services"
)

type LibraryController struct {
	LibraryManager services.LibraryManager
}

func (controller *LibraryController) StartLibraryController() {
	//Displaying the menu
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose an option:")
		fmt.Println("1. Add a book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Display available books")
		fmt.Println("4. Borrow a book")
		fmt.Println("5. Return a book")
		fmt.Println("6. Display borrowed books")
		fmt.Println("7. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			controller.AddBook()
		case 2:
			controller.RemoveBook()
		case 3:
			controller.ListAvailableBooks()
		case 4:
			controller.BorrowBook()
		case 5:
			controller.ReturnBook()
		case 6:
			controller.ListBorrowedBooks()
		case 7:
			return
		default:
			fmt.Println("Invalid choice")

		}
		fmt.Println("Do you want to continue? (yes/no)")
		scanner.Scan()
		if strings.ToLower(scanner.Text()) != "yes" {
			break
		}
	}

}

var nextID int = 1

func (controller *LibraryController) AddBook() {
	scanner := bufio.NewScanner(os.Stdin)
	var title string
	var author string

	fmt.Println("Enter the book title:")
	scanner.Scan()
	title = scanner.Text()
	fmt.Println("Enter the book author:")
	scanner.Scan()
	author = scanner.Text()

	newBook := models.Book{ID: nextID, Title: title, Author: author, Status: "Available"}
	nextID++
	controller.LibraryManager.AddBook(&newBook)
	fmt.Println("Book added successfully")
}

func (controller *LibraryController) RemoveBook() {
	var bookID int
	fmt.Println("Enter the book ID:")
	fmt.Scanln(&bookID)
	if reflect.TypeOf(bookID).Kind() != reflect.Int {
		fmt.Println("Invalid book ID. Please enter a valid number.")
		return
	}
	err := controller.LibraryManager.RemoveBook(bookID)
	if err != nil {
		fmt.Println("Seems like there is no book with that ID")
	} else {
		fmt.Println("Book removed successfully")
	}
}

func (controller *LibraryController) ListAvailableBooks() {
	availableBooks := controller.LibraryManager.ListAvailableBooks()
	if len(availableBooks) == 0 {
		fmt.Println("No available books")
	}
	for _, book := range availableBooks {
		fmt.Printf("ID: %d\n", book.ID)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Println("Status: Available")
	}
}
func (controller *LibraryController) BorrowBook() {
	var bookID int
	var memberID int
	fmt.Println("Enter the book ID:")
	fmt.Scanln(&bookID)
	fmt.Println("Enter your(the member) ID:")
	fmt.Scanln(&memberID)
	if reflect.TypeOf(bookID).Kind() != reflect.Int || reflect.TypeOf(memberID).Kind() != reflect.Int {
		fmt.Println("Invalid book ID. Please enter a valid number.")
		return
	}
	err := controller.LibraryManager.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Seems like there is no available book or member with that ID")
	} else {
		fmt.Println("Book borrowed successfully")
	}
}

func (controller *LibraryController) ReturnBook() {
	var bookID int
	var memberID int
	fmt.Println("Enter the book ID:")
	fmt.Scanln(&bookID)
	fmt.Println("Enter your(the member) ID:")
	fmt.Scanln(&memberID)
	// Check if bookID is a number
	if reflect.TypeOf(bookID).Kind() != reflect.Int || reflect.TypeOf(memberID).Kind() != reflect.Int {
		fmt.Println("Invalid book ID. Please enter a valid number.")
		return
	}
	if err := controller.LibraryManager.ReturnBook(bookID, memberID); err != nil {
		fmt.Println("Seems like there is no borrowed book or member with that ID")
	} else {
		fmt.Println("Book returned successfully")
	}
}

func (controller *LibraryController) ListBorrowedBooks() {
	var memberID int
	fmt.Println("Enter your(the member) ID:")
	fmt.Scanln(&memberID)
	if reflect.TypeOf(memberID).Kind() != reflect.Int {
		fmt.Println("Invalid book ID. Please enter a valid number.")
		return
	}
	borrowedBooks, err := controller.LibraryManager.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Println("Seems like there is no member with that ID")
	} else {
		if len(borrowedBooks) == 0 {
			fmt.Println("No borrowed books")
		}
		for _, book := range borrowedBooks {
			fmt.Printf("ID: %d\n", book.ID)
			fmt.Printf("Title: %s\n", book.Title)
			fmt.Printf("Author: %s\n", book.Author)
			fmt.Println("Status: Borrowed")
		}
	}
}
