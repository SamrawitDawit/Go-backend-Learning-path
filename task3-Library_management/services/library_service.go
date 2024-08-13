package services

import (
	"errors"
	"task3-Library_management/models"
)

type Library struct {
	Books   map[int]*models.Book
	Members map[int]*models.Member
}

type LibraryManager interface {
	AddBook(book *models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []*models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

func NewLibrary(books map[int]*models.Book, members map[int]*models.Member) LibraryManager {
	return &Library{
		Books:   books,
		Members: members,
	}
}
func (library *Library) AddBook(book *models.Book) {
	library.Books[book.ID] = book
}
func (library *Library) RemoveBook(bookID int) error {
	if library.Books[bookID] == nil {
		return errors.New("book not found")
	}
	delete(library.Books, bookID)
	return nil
}

func (library *Library) BorrowBook(bookID int, memberID int) error {
	member := library.Members[memberID]
	if member == nil {
		return errors.New("member not found")
	}
	if book, ok := library.Books[bookID]; ok {
		if book.Status == "Available" {
			book.Status = "Borrowed"
			member.BorrowedBooks = append(member.BorrowedBooks, *book)
			return nil
		}
	}
	return errors.New("book is not available")
}

func (library *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := library.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}
	for i, book := range member.BorrowedBooks {
		if book.ID == bookID {
			library.Books[bookID].Status = "Available"
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			return nil
		}
	}
	return errors.New("you haven't borrowed such book")
}

func (library *Library) ListAvailableBooks() []*models.Book {
	var availableBooks []*models.Book
	for _, book := range library.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (library *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member := library.Members[memberID]
	if member == nil {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}
