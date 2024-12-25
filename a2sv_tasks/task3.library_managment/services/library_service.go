package services

import (
	"errors"
	"library_manager/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	AddMember(member models.Member)
	RemoveBook(bookID int) error
	BorrowBook(bookID, memberID int) error
	ReturnBook(bookID, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memeberID int) []models.Book
}

type LibraryService struct {
	BookID      int
	MemberID    int
	BooksList   map[int]models.Book
	MembersList map[int]models.Member
	BookStatus  []string
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		BookID:      -1,
		MemberID:    -1,
		BooksList:   make(map[int]models.Book),
		MembersList: make(map[int]models.Member),
		BookStatus:  []string{"available", "borrowed"},
	}
}

func (ls *LibraryService) AddBook(book models.Book) {
	ls.BookID++
	book.ID = ls.BookID
	ls.BooksList[ls.BookID] = book
}

func (ls *LibraryService) AddMember(member models.Member) {
	ls.MemberID++
	member.ID = ls.MemberID
	ls.MembersList[ls.MemberID] = member
}

func (ls *LibraryService) RemoveBook(id int) error {
	if _, ok := ls.BooksList[id]; !ok {
		return errors.New("no book found with the given id")
	}
	delete(ls.BooksList, id)
	return nil
}

func (ls *LibraryService) BorrowBook(bookID, memberID int) error {
	if _, ok := ls.BooksList[bookID]; !ok {
		return errors.New("no book found with the given id")
	}
	if _, ok := ls.MembersList[memberID]; !ok {
		return errors.New("no member found with the given id")
	}

	if ls.BooksList[bookID].Status == ls.BookStatus[1] {
		return errors.New("book is already borrowed")
	}

	// add the book to the list of member borrowed list and change the book status
	book := ls.BooksList[bookID]
	book.Status = ls.BookStatus[1]
	ls.BooksList[bookID] = book

	member := ls.MembersList[memberID]
	member.BorrowedBooks[bookID] = book
	ls.MembersList[memberID] = member

	return nil

}

func (ls *LibraryService) ReturnBook(bookID, memeberID int) error {
	if _, exist := ls.BooksList[bookID]; !exist {
		return errors.New("book does not exists")
	}
	if _, exist := ls.MembersList[memeberID]; !exist {
		return errors.New("member does not exists")
	}

	if _, exist := ls.MembersList[memeberID].BorrowedBooks[bookID]; !exist {
		return errors.New("you didn't borrowed this book")
	}

	// Remove the book from the MembersList borrowed book list and change the book status
	delete(ls.MembersList[memeberID].BorrowedBooks, bookID)
	book := ls.BooksList[bookID]
	book.Status = ls.BookStatus[0]
	ls.BooksList[bookID] = book

	return nil
}
