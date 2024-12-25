package controllers

import (
	"bufio"
	"fmt"
	"library_manager/models"
	"library_manager/services"
	"os"
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

func (lc *LibraryController) AddBook() {
	fmt.Print("Book Title: ")
	lc.scannerInstance.Scan()
	title := lc.scannerInstance.Text()
	fmt.Print("Book Author: ")
	lc.scannerInstance.Scan()
	author := lc.scannerInstance.Text()
	status := lc.LibraryServiceInstance.BookStatus[0]

	book := models.Book{
		ID:     0,
		Title:  title,
		Author: author,
		Status: status,
	}

	lc.LibraryServiceInstance.AddBook(book)

}
