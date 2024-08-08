package models

type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}

var books = []Book{}
