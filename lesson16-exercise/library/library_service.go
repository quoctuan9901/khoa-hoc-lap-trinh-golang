package library

import (
	"fmt"

	"quoctuan.com/hoc-golang/utils"
)

func AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.GetNonEmptyString("Nhap tieu de: ")
	author := utils.GetNonEmptyString("Nhap tac gia: ")

	if err := lib.AddBookStore(id, title, author); err != nil {
		return err
	}

	fmt.Printf("✅ Them sach thanh cong! ID: %s", id)

	return nil
}

func ListBooks(lib *Library) error {
	books := lib.ListBooksStore()
	if len(books) == 0 {
		fmt.Println("Thu vien hien tai chua co sach. Vui long them sach")
		return nil
	}

	for _, book := range books {
		status := "Con"
		if book.IsBorrowed {
			status = "Da muon"
		}

		fmt.Printf("Id: %s, Tieu de: %s, Tac gia: %s, Trang thai: %s \n", book.Id, book.Title, book.Author, status)
	}

	return nil
}

func AddBorrower(lib *Library) error {
	id := utils.GenerateId()
	name := utils.GetNonEmptyString("Nhap ten: ")
	email := utils.GetNonEmptyString("Nhap email: ")

	if err := lib.AddBorrowerStore(id, name, email); err != nil {
		return err
	}

	fmt.Printf("✅ Them nguoi muon sach thanh cong! ID: %s", id)

	return nil
}

func ListBorrowers(lib *Library) error {
	borrowers := lib.ListBorrowersStore()
	if len(borrowers) == 0 {
		fmt.Println("Thu vien hien chua co nguoi nao dang ky muon sach")
		return nil
	}

	fmt.Println("Danh sach nguoi muon:")
	for _, borrower := range borrowers {
		fmt.Printf("Id: %s, Ten: %s, Email: %s \n", borrower.Id, borrower.Name, borrower.Email)
	}

	return nil
}

func BorrowBook(lib *Library) error {
	transactionId := utils.GenerateId()
	bookId := utils.GetNonEmptyString("Nhap ID sach: ")
	borrowerId := utils.GetNonEmptyString("Nhap ID nguoi muon sach: ")

	if err := lib.BorrowBookStore(transactionId, bookId, borrowerId); err != nil {
		return err
	}

	fmt.Printf("✅ Muon sach thanh cong! ID giao dich: %s", transactionId)

	return nil
}

func ListBorrowHistory(lib *Library) error {
	borrowerId := utils.GetNonEmptyString("Nhap ID nguoi muon: ")
	transactions := lib.ListBorrowHistoryByBorrowerStore(borrowerId)
	if len(transactions) == 0 {
		fmt.Println("Khong co lich su mon sach voi ID nay")
		return nil
	}

	fmt.Println("Lich su muon sach: ")
	for _, transaction := range transactions {
		returnDate := "Chua tra"
		if !transaction.ReturnDate.IsZero() {
			returnDate = transaction.ReturnDate.Format("2006-01-02")
		}
		fmt.Printf("Giao dich %s, Sach: %s, Ngay muon: %s, Ngay tra: %s \n", transaction.TransactionId, lib.GetBookTitleStore(transaction.BookId), transaction.BorrowDate.Format("2006-01-02"), returnDate)
	}


	return nil
}

func ReturnBook(lib *Library) error {
	transactionId := utils.GetNonEmptyString("Nhap Id giao dich: ")
	if err := lib.ReturnBookStore(transactionId); err != nil {
		return err
	}

	fmt.Println("✅ Tra sach thanh cong!")
	return nil
}

func SearchBooks(lib *Library) error {
	query := utils.GetNonEmptyString("Nhap tieu de hoac tac gia de tim kiem: ")

	books := lib.SearchBooksStore(query)
	if len(books) == 0 {
		fmt.Println("Khong tim thay sach nao trong thu vien voi noi dung nay.")
		return nil
	}

	for _, book := range books {
		status := "Con"
		if book.IsBorrowed {
			status = "Da muon"
		}

		fmt.Printf("Id: %s, Tieu de: %s, Tac gia: %s, Trang thai: %s \n", book.Id, book.Title, book.Author, status)
	}

	return nil
}
