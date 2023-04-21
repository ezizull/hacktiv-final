package book

import (
	"time"

	"gorm.io/gorm"
)

// Book is a struct that contains the book model
type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" example:"book title"`
	UserID      int       `json:"user_id" gorm:"index"`
	Description string    `json:"description" example:"book description"`
	CreatedAt   time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt   *gorm.DeletedAt
}

// TableName overrides the table name used by User to `users`
func (*Book) TableName() string {
	return "books"
}
