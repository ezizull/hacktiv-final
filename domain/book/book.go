package book

import (
	"time"
)

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" example:"book title"`
	UserID      int    `json:"user_id" gorm:"index"`
	Description string `json:"description" example:"book description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (*Book) TableName() string {
	return "books"
}

// PaginationResultBook is a struct that contains the pagination result for book
type PaginationResultBook struct {
	Data       *[]Book
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
