package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task3-Library_management/models"
	"task3-Library_management/services"
)

var LibraryController services.Library

func StartLibraryController() {
	//Displaying the menu
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
		AddBook()
	case 2:
		RemoveBook()
	case 3:
		ListAvailableBooks()
	case 4:
		BorrowBook()
	case 5:
		ReturnBook()
	case 6:
		ListBorrowedBooks()
	case 7:
		return
	default:
		fmt.Println("Invalid choice")

	}
}

var nextID int = 1

func AddBook() {
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
	LibraryController.AddBook(newBook)
	fmt.Println("Book added successfully")
	fmt.Println("Do you want to do another operation? (yes/no)")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "yes" {
		StartLibraryController()
	}
}

func RemoveBook() {
	scanner := bufio.NewScanner(os.Stdin)
	var bookID int
	fmt.Println("Enter the book ID:")
	fmt.Scanln(&bookID)
	err := LibraryController.RemoveBook(bookID)
	if err != nil {
		fmt.Println("Seems like there is no book with that ID")
	} else {
		fmt.Println("Book removed successfully")
	}
	fmt.Println("Do you want to do another operation? (yes/no)")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "yes" {
		StartLibraryController()
	}
}

func ListAvailableBooks() {
	scanner := bufio.NewScanner(os.Stdin)
	availableBooks := LibraryController.ListAvailableBooks()
	if len(availableBooks) == 0 {
		fmt.Println("No available books")
	}
	for _, book := range availableBooks {
		fmt.Printf("ID: %d\n", book.ID)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Println("Status: Available")
	}
	fmt.Println("Do you want to do another operation? (yes/no)")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "yes" {
		StartLibraryController()
	}
}
func BorrowBook() {
	scanner := bufio.NewScanner(os.Stdin)
	var bookID int
	var memberID int
	fmt.Println("Enter the book ID:")
	fmt.Scanln(&bookID)
	fmt.Println("Enter your(the member) ID:")
	fmt.Scanln(&memberID)
	err := LibraryController.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Seems like there is no available book or member with that ID")
	} else {
		fmt.Println("Book borrowed successfully")
	}
	fmt.Println("Do you want to do another operation? (yes/no)")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "yes" {
		StartLibraryController()
	}
}

func ReturnBook() {
	scanner := bufio.NewScanner(os.Stdin)
	var bookID int
	var memberID int
	fmt.Println("Enter the book ID:")
	fmt.Scanln(&bookID)
	fmt.Println("Enter your(the member) ID:")
	fmt.Scanln(&memberID)
	err := LibraryController.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Seems like there is no borrowed book or member with that ID")
	} else {
		fmt.Println("Book returned successfully")
	}
	fmt.Println("Do you want to do another operation? (yes/no)")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "yes" {
		StartLibraryController()
	}
}

func ListBorrowedBooks() {
	scanner := bufio.NewScanner(os.Stdin)
	var memberID int
	fmt.Println("Enter your(the member) ID:")
	fmt.Scanln(&memberID)
	borrowedBooks, err := LibraryController.ListBorrowedBooks(memberID)
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

	fmt.Println("Do you want to do another operation? (yes/no)")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "yes" {
		StartLibraryController()
	}
}
