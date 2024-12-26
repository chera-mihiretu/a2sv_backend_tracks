package main

import (
	"fmt"
	"library_manager/controllers"
)

func main() {
	fmt.Println("Welcome To Library Management")
	libraryController := controllers.NewLibraryController()
	userID := libraryController.AddMember()
	for {
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		choice := libraryController.TakeInput("Enter your choice: ")
		switch choice {
		case "1":
			libraryController.AddBook()
		case "2":
			libraryController.RemoveBook()
		case "3":
			libraryController.BorrowBook(userID)
		case "4":
			libraryController.ReturnBook(userID)
		case "5":
			libraryController.ListBorrowedBooks(userID)
		default:
			fmt.Println("Invalid Input")
		}
	}
}
