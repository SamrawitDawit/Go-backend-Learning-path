package models

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

var Members = []Member{
	{ID: 1, Name: "Samrawit Dawit", BorrowedBooks: []Book{}},
	{ID: 2, Name: "Rediet Woudema", BorrowedBooks: []Book{}},
}
