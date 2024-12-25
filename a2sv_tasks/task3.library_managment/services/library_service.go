package services

import (
	"errors"
	"library_manager/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID, memberID int) error
	ReturnBook(bookID, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memeberID int) []models.Book
}

type LibraryService struct {
	bookID   int
	memberID int
	books    map[int]models.Book
	members  map[int]models.Member
}

func NewLibraryService() LibraryService {
	return LibraryService{
		bookID:   0,
		memberID: 0,
		books:    make(map[int]models.Book),
		members:  make(map[int]models.Member),
	}
}

func (ls *LibraryService) GetBookMeId() int {
	return ls.bookID + 1
}

func (ls *LibraryService) GetMemberId() int {
	return ls.memberID + 1
}

func (ls *LibraryService) AddBook(book models.Book) {
	ls.bookID++
	book.ID = ls.bookID
	ls.books[ls.bookID] = book
}

func (ls *LibraryService) RemoveBook(id int) error {
	if _, ok := ls.books[id]; !ok {
		return errors.New("no book found with the given id")
	}
	delete(ls.books, id)
	return nil
}
