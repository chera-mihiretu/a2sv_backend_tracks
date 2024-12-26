package controllers

import (
	"bufio"
	"fmt"
	"library_manager/models"
	"library_manager/services"
	"os"
	"strconv"
)

type LibraryController struct {
	LibraryServiceInstance *services.LibraryService
	scannerInstance        *bufio.Scanner
}

func NewLibraryController() *LibraryController {
	return &LibraryController{
		LibraryServiceInstance: services.NewLibraryService(),
		scannerInstance:        bufio.NewScanner(bufio.NewReader(os.Stdin)),
	}

}

func (lc *LibraryController) TakeInput(input string) string {
	fmt.Print(input)
	lc.scannerInstance.Scan()

	variable := lc.scannerInstance.Text()

	return variable

}

func (lc *LibraryController) AddBook() {

	title := lc.TakeInput("Book Title: ")
	author := lc.TakeInput("Book Author: ")
	status := lc.LibraryServiceInstance.BookStatus[0]

	book := models.Book{
		ID:     0,
		Title:  title,
		Author: author,
		Status: status,
	}

	lc.LibraryServiceInstance.AddBook(book)

}

func (lc *LibraryController) AddMember() int {

	name := lc.TakeInput("Enter User Name: ")

	member := models.Member{
		ID:            0,
		Name:          name,
		BorrowedBooks: make(map[int]bool),
	}

	id := lc.LibraryServiceInstance.AddMember(member)
	return id

}

func (lc *LibraryController) RemoveBook() {
	var bookID int
	var err error
	for {
		id := lc.TakeInput("Enter The Book Id")
		bookID, err = strconv.Atoi(id)
		if err == nil {
			break
		}
	}

	removeError := lc.LibraryServiceInstance.RemoveBook(bookID)
	if removeError != nil {
		fmt.Println("Look like the book you entered Does not exists")
		return
	}

	fmt.Println("Book removed successfully")
}

func (lc *LibraryController) BorrowBook(memberID int) {
	var bookID int
	var err error
	for {
		id := lc.TakeInput("Enter The Book Id")
		bookID, err = strconv.Atoi(id)
		if err == nil {
			break
		}
	}
	borrowError := lc.LibraryServiceInstance.BorrowBook(bookID, memberID)

	if borrowError != nil {
		fmt.Println("There ")
	}
}

func (lc *LibraryController) ReturnBook(memberID int) {

	var bookID int
	var err error
	for {
		id := lc.TakeInput("Enter The Book Id")
		bookID, err = strconv.Atoi(id)
		if err == nil {
			break
		}
	}

	returnError := lc.LibraryServiceInstance.ReturnBook(bookID, memberID)

	if returnError != nil {
		fmt.Println("You may entered wrong book id")
		return
	}

	fmt.Println("Book returned")

}

func (lc *LibraryController) ListAvailableBooks() {
	books := lc.LibraryServiceInstance.ListAvailableBooks()
	fmt.Printf("%-5s %-20s %-20s %-20s", "ID", "TITLE", "AUTHOR", "STATUS")
	for id, book := range books {
		fmt.Printf("%-5d %-20s %-20s %-20s", id, book.Title, book.Author, book.Status)
	}
}

func (lc *LibraryController) ListBorrowedBooks(memberID int) {
	books := lc.LibraryServiceInstance.ListBorrowedBooks(memberID)
	fmt.Printf("%-5s %-20s %-20s %-20s", "ID", "TITLE", "AUTHOR", "STATUS")
	for id, book := range books {
		fmt.Printf("%-5d %-20s %-20s %-20s", id, book.Title, book.Author, book.Status)
	}
}
