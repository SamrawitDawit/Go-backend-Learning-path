package services

import (
	"errors"
	"task3-Library_management/models"
)

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func (library *Library) AddBook(book models.Book) {
	library.Books[book.ID] = book
}
func (library *Library) RemoveBook(bookID int) {
	delete(library.Books, bookID)
}

func (library *Library) BorrowBook(bookID int, memberID int) error {
	if book, ok := library.Books[bookID]; ok {
		if book.Status == "Available" {
			book.Status = "Borrowed"
			member := library.Members[memberID]
			member.BorrowedBooks = append(member.BorrowedBooks, book)
			return nil
		}
	}
	return errors.New("book is not available")
}

func (library *Library) ReturnBook(bookID int, memberID int) error {
	member := library.Members[memberID]
	for i, book := range member.BorrowedBooks {
		if book.ID == bookID {
			book.Status = "Available"
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			return nil
		}
	}
	return errors.New("you haven't borrowed such book")
}

func (library *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range library.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (library *Library) ListBorrowedBooks(memberID int) []models.Book {
	return library.Members[memberID].BorrowedBooks
}
