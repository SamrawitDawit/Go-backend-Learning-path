package controllers

import (
	"task3-Library_management/models"
	"task3-Library_management/services"
)

type LibraryController struct {
	LibraryService services.LibraryManager
}

func (controller *LibraryController) AddBook(book models.Book) {
	controller.LibraryService.AddBook(book)
}
func (controller *LibraryController) RemoveBook(bookID int) {
	controller.LibraryService.RemoveBook(bookID)
}
func (controller *LibraryController) BorrowBook(bookID int, memberID int) {
	controller.LibraryService.BorrowBook(bookID, memberID)
}
func (controller *LibraryController) ReturnBook(bookID int, memberID int) {
	controller.LibraryService.ReturnBook(bookID, memberID)
}
func (controller *LibraryController) ListAvailableBooks() []models.Book {
	res := controller.LibraryService.ListAvailableBooks()
	return res
}
func (controller *LibraryController) ListBorrowedBooks(memberID int) []models.Book {
	res := controller.LibraryService.ListBorrowedBooks(memberID)
	return res
}