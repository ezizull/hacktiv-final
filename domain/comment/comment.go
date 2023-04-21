package comment

import "time"

// Comment is a struct that contains the comment information
type Comment struct {
	ID        int       `json:"id" example:"1099" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"index"`
	PhotoID   int       `json:"photo_id" gorm:"index"`
	Message   string    `json:"message" example:"caption"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt time.Time
}

// TableName overrides the table name used by Comment to `comments`
func (*Comment) TableName() string {
	return "comments"
}
