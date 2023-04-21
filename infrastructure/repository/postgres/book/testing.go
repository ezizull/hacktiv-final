package book

import bookDomain "hacktiv/final-project/domain/book"

type BookTesting interface {
	GetAll(page int64, limit int64) (*bookDomain.PaginationResultBook, error)
	UserGetAll(page int64, userId int, limit int64) (*bookDomain.PaginationResultBook, error)
	Create(newBook *bookDomain.Book) (createdBook *bookDomain.Book, err error)
	GetByID(id int) (*bookDomain.Book, error)
	UserGetByID(id int, userId int) (*bookDomain.Book, error)
	GetOneByMap(bookMap map[string]interface{}) (*bookDomain.Book, error)
	Update(id uint, bookMap map[string]interface{}) (*bookDomain.Book, error)
	Delete(id int) (err error)
}
