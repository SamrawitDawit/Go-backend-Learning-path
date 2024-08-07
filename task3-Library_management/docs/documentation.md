#Library Management System

##Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Project Structure](#project-structure)
6. [API Endpoints](#api-endpoints)
7. [Contributing](#contributing)

## Introduction
The Library Management System is designed to help manage books and members in a library. It allows adding, removing, borrowing, and returning books, as well as listing available and borrowed books.


## Features
- Add books to the library
- Remove books from the library
- Borrow books by members
- Return borrowed books
- List all available books
- List all borrowed books by a specific member

## Installation

### Prerequisites
- Go 1.16 or higher

### Steps
1. Clone the repository:
```sh
git clone https://github.com/SamrawitDawit/Go-backend-Learning-path.git
```
2. Navigate to the project directory:
```sh
cd task3-Library_management
```
3. Build the project:
```sh
go build
```

## Usage

To run the application:
```sh
go run main.go
```

## Project Structure
Library Management
├── controllers
│   └── library_controller.go
├── models
│   └── book.go
│   └── member.go
├── services
│   └── library_service.go
├── main.go
├── go.mod
└── docs
    └──documentation.md

## API Endpoints

### Add a book
- adds a book to the library.
   Parameters:
    book: The book to be added.
   Returns:
    None.
### Remove a book 
- removes a book from the library.
   Parameters:
    bookID: The ID of the book to be deleted
   Returns:
    None.
### Borrow a book 
-  borrows a book from the library for a specific member.
   Parameters: 
    bookID and memberID
   Returns:
    None.
### Return a book
-  returns a book to the library.
   Parameters:
    bookID and memberID
   Returns:
    None.
### List Available Books
- lists the available books in the library.
  Parameters:
    None.
  Returns: 
    A list of available books
### List Borrowed Books
- lists the books that a specific member borrowed.
  Parameters:
    memberID
  Returns: 
    A list of borrowed books the memeber borrowed.

## Contributing
1. Fork the repository
2. Create a new branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull request












