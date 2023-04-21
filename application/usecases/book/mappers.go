package book

// Package medicine provides the use case for medicine

import (
	bookDomain "hacktiv/final-project/domain/book"
)

func (n *NewBook) toDomainMapper() *bookDomain.Book {
	return &bookDomain.Book{
		Title:       n.Title,
		UserID:      n.UserID,
		Description: n.Description,
	}
}
