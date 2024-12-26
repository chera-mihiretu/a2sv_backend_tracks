package test

import (
	"library_manager/models"
	"library_manager/services"
	"testing"
)

func TestLibraryService(t *testing.T) {

	libraryService := services.NewLibraryService()

	test_cases := map[int]models.Book{
		0: {
			ID:     0,
			Title:  "The Leaves",
			Author: "Chera Mihiretu",
			Status: libraryService.BookStatus[0],
		},
		1: {
			ID:     1,
			Title:  "The Sun",
			Author: "Chera Mihiretu",
			Status: libraryService.BookStatus[0],
		},
	}

	for _, book := range test_cases {
		libraryService.AddBook(book)
	}

	t.Log(libraryService.BooksList)

	for id, book := range test_cases {
		if libraryService.BooksList[id] != test_cases[id] {
			t.Errorf("Expected %v, but got %v", book, libraryService.BooksList[id])
		}
	}

}

func TestLibraryService_AddMember(t *testing.T) {
	libraryService := services.NewLibraryService()

	test_cases := map[int]models.Member{
		0: {
			ID:            0,
			Name:          "Chera Mihiretu",
			BorrowedBooks: make(map[int]bool),
		},
		1: {
			ID:            1,
			Name:          "Chera Mihiretu",
			BorrowedBooks: make(map[int]bool),
		},
	}

	for _, member := range test_cases {
		libraryService.AddMember(member)
	}

	t.Log(libraryService.MembersList)

	for id, member := range test_cases {
		if libraryService.MembersList[id].Name != test_cases[id].Name && libraryService.MembersList[id].ID != test_cases[id].ID {
			t.Errorf("Expected %v, but got %v", member, libraryService.MembersList[id])
		}
	}
}

func TestLibraryService_RemoveBook(t *testing.T) {
	libraryService := services.NewLibraryService()

	test_cases := map[int]models.Book{
		0: {
			ID:     0,
			Title:  "The Leaves",
			Author: "Chera Mihiretu",
			Status: libraryService.BookStatus[0],
		},
		1: {
			ID:     1,
			Title:  "The Sun",
			Author: "Chera Mihiretu",
			Status: libraryService.BookStatus[0],
		},
	}

	for _, book := range test_cases {
		libraryService.AddBook(book)
	}

	err := libraryService.RemoveBook(2)
	if err == nil {
		t.Errorf("Expected error for non-existent book ID, but got none")
	}

	err = libraryService.RemoveBook(0)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if _, exists := libraryService.BooksList[0]; exists {
		t.Errorf("Expected book to be removed, but it still exists")
	}
}

func TestLibraryService_BorrowBook(t *testing.T) {
	libraryService := services.NewLibraryService()

	book := models.Book{
		ID:     0,
		Title:  "The Leaves",
		Author: "Chera Mihiretu",
		Status: libraryService.BookStatus[0],
	}
	libraryService.AddBook(book)

	member := models.Member{
		ID:            0,
		Name:          "Chera Mihiretu",
		BorrowedBooks: make(map[int]bool),
	}
	libraryService.AddMember(member)

	err := libraryService.BorrowBook(1, 0)
	if err == nil {
		t.Errorf("Expected error for non-existent book ID, but got none")
	}

	err = libraryService.BorrowBook(0, 1)
	if err == nil {
		t.Errorf("Expected error for non-existent member ID, but got none")
	}

	err = libraryService.BorrowBook(0, 0)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if libraryService.BooksList[0].Status != libraryService.BookStatus[1] {
		t.Errorf("Expected book status to be 'borrowed', but got %v", libraryService.BooksList[0].Status)
	}
	if libraryService.MembersList[0].BorrowedBooks[0] == false {
		t.Errorf("Expected member to have borrowed book, but got none")
	}
}

func TestLibraryService_ReturnBook(t *testing.T) {
	libraryService := services.NewLibraryService()

	book := models.Book{
		ID:     0,
		Title:  "The Leaves",
		Author: "Chera Mihiretu",
		Status: libraryService.BookStatus[0],
	}
	libraryService.AddBook(book)

	member := models.Member{
		ID:            0,
		Name:          "Chera Mihiretu",
		BorrowedBooks: make(map[int]bool),
	}
	libraryService.AddMember(member)

	libraryService.BorrowBook(0, 0)

	err := libraryService.ReturnBook(1, 0)
	if err == nil {
		t.Errorf("Expected error for non-existent book ID, but got none")
	}

	err = libraryService.ReturnBook(0, 1)
	if err == nil {
		t.Errorf("Expected error for non-existent member ID, but got none")
	}

	err = libraryService.ReturnBook(0, 0)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if libraryService.BooksList[0].Status != libraryService.BookStatus[0] {
		t.Errorf("Expected book status to be 'available', but got %v", libraryService.BooksList[0].Status)
	}
	if _, exists := libraryService.MembersList[0].BorrowedBooks[0]; exists {
		t.Errorf("Expected member to have returned book, but still has it")
	}
}
