package main

import (
	"fmt"
	"library_manager/controllers"
	"os"
	"os/exec"
)

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

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
		ClearTerminal()
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
			libraryController.ListAvailableBooks()
		case "6":
			libraryController.ListBorrowedBooks(userID)
		case "7":

		default:
			fmt.Println("Invalid Input")
		}

	}
}
