package book

import (
	bookDomain "hacktiv/final-project/domain/book"
	bookRepository "hacktiv/final-project/infrastructure/repository/postgres/book"
)

type BookTesting interface {
	GetAll(page int64, limit int64) (*bookDomain.PaginationResultBook, error)
	UserGetAll(page int64, userId int, limit int64) (*bookDomain.PaginationResultBook, error)
	GetByID(id int) (*bookDomain.Book, error)
	UserGetByID(id int, userId int) (*bookDomain.Book, error)
	Create(book *NewBook) (*bookDomain.Book, error)
	GetByMap(medicineMap map[string]interface{}) (*bookDomain.Book, error)
	Delete(id int) (err error)
	Update(id uint, medicineMap map[string]interface{}) (*bookDomain.Book, error)
}

func NewTesting(bookTest bookRepository.BookTesting) BookTesting {
	return &Service{
		bookTesting: bookTest,
	}
}
