package book

import bookDomain "hacktiv/final-project/domain/book"

func (book *Book) toDomainMapper() *bookDomain.Book {
	return &bookDomain.Book{
		ID:          book.ID,
		Title:       book.Title,
		UserID:      book.UserID,
		Description: book.Description,
		CreatedAt:   book.CreatedAt,
	}
}

func fromDomainMapper(book *bookDomain.Book) *Book {
	return &Book{
		ID:          book.ID,
		Title:       book.Title,
		UserID:      book.UserID,
		Description: book.Description,
		CreatedAt:   book.CreatedAt,
	}
}

func arrayToDomainMapper(books *[]Book) *[]bookDomain.Book {
	booksDomain := make([]bookDomain.Book, len(*books))
	for i, book := range *books {
		booksDomain[i] = *book.toDomainMapper()
	}

	return &booksDomain
}
