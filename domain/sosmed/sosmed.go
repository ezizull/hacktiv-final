package sosmed

import "time"

// SocialMedia is a struct that contains the social media information
type SocialMedia struct {
	ID             int       `json:"id" example:"1099" gorm:"primaryKey"`
	Name           string    `json:"name" example:"caption"`
	SocialMediaUrl string    `json:"social_media_url" example:"www.sosmed.com"`
	UserID         int       `json:"user_id" gorm:"index"`
	CreatedAt      time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt      time.Time
}

// TableName overrides the table name used by SocialMedia to `social_media`
func (*SocialMedia) TableName() string {
	return "social_media"
}
