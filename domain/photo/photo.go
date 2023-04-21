package photo

import "time"

// Photo is a struct that contains the photo information
type Photo struct {
	ID        int       `json:"id" example:"1099" gorm:"primaryKey"`
	Title     string    `json:"title" example:"title"`
	Caption   string    `json:"caption" example:"caption"`
	PhotoUrl  string    `json:"photo_url" example:"www.photo.com"`
	UserID    int       `json:"user_id" gorm:"index"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt time.Time
}

// TableName overrides the table name used by User to `users`
func (*Photo) TableName() string {
	return "photos"
}
