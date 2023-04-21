package book

// NewBook is a struct that contains the data for a new book
type NewBook struct {
	Title       string `json:"title" example:"book title"`
	UserID      int    `json:"user_id" gorm:"index"`
	Description string `json:"description" example:"book description"`
}
