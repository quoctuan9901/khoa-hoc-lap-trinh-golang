package library

import (
	"fmt"
	"strings"
	"time"

	"quoctuan.com/hoc-golang/models"
)

type Library struct {
	books        map[string]models.Book
	borrowers    map[string]models.Borrower
	transactions map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		books:        make(map[string]models.Book),
		borrowers:    make(map[string]models.Borrower),
		transactions: make(map[string]models.Transaction),
	}
}

func (lib *Library) AddBookStore(id, title, author string) error {
	if _, exists := lib.books[id]; exists {
		return fmt.Errorf("sach voi Id %s da ton tai", id)
	}

	lib.books[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}

	return nil
}

func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.books))
	for _, book := range lib.books {
		books = append(books, book)
	}

	return books
}

func (lib *Library) AddBorrowerStore(id, name, email string) error {
	if _, exists := lib.borrowers[id]; exists {
		return fmt.Errorf("nguoi muon voi Id %s da ton tai", id)
	}

	lib.borrowers[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return nil
}

func (lib *Library) ListBorrowersStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrowers))
	for _, borrower := range lib.borrowers {
		borrowers = append(borrowers, borrower)
	}

	return borrowers
}

func (lib *Library) BorrowBookStore(id, bookId, borrowerId string) error {
	book, bookExists := lib.books[bookId]
	if !bookExists {
		return fmt.Errorf("sach voi ID %s khong ton tai", bookId)
	}

	if book.IsBorrowed {
		return fmt.Errorf("sach %s da duoc muon", book.Title)
	}

	if _, borrowerExists := lib.borrowers[borrowerId]; !borrowerExists {
		return fmt.Errorf("nguoi muon sach voi ID %s khong ton tai", borrowerId)
	}

	if _, transactionExists := lib.transactions[id]; transactionExists {
		return fmt.Errorf("giao dich voi ID %s da ton tai", id)
	}

	book.IsBorrowed = true
	lib.books[bookId] = book
	lib.transactions[id] = models.Transaction{
		TransactionId: id,
		BookId:        bookId,
		BorrowerId:    borrowerId,
		BorrowDate:    time.Now(),
	}

	return nil
}

func (lib *Library) ListBorrowHistoryByBorrowerStore(borrowerId string) []models.Transaction {
	if _, borrowerExists := lib.borrowers[borrowerId]; !borrowerExists {
		return nil
	}

	var history []models.Transaction
	for _, transaction := range lib.transactions {
		if transaction.BorrowerId == borrowerId {
			history = append(history, transaction)
		}
	}

	return history
}

func (lib *Library) GetBookTitleStore(bookId string) string {
	book := lib.books[bookId]

	return book.Title
}

func (lib *Library) ReturnBookStore(transactionId string) error {
	transaction, exists := lib.transactions[transactionId]
	if !exists {
		return fmt.Errorf("giao dich voi ID %s khong ton tai", transactionId)
	}

	if !transaction.ReturnDate.IsZero() {
		return fmt.Errorf("sach trong giao dich %s da duoc tra roi", transactionId)
	}

	// Update IsBorrowed of book
	book, bookExists := lib.books[transaction.BookId]
	if !bookExists {
		return fmt.Errorf("sach voi ID %s khong ton tai", transaction.BookId)
	}
	book.IsBorrowed = false
	lib.books[transaction.BookId] = book

	// Update ReturnDate of transaction
	transaction.ReturnDate = time.Now()
	lib.transactions[transactionId] = transaction
	return nil
}

func (lib *Library) SearchBooksStore(query string) []models.Book {
	query = strings.ToLower(query)

	var results []models.Book
	for _, book := range lib.books {
		if strings.Contains(strings.ToLower(book.Title), query) || strings.Contains(strings.ToLower(book.Author), query) {
			results = append(results, book)
		}
	}

	return results
}