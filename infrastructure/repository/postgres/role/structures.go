package role

import "time"

// Role is a struct that contains the role information
type Role struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name used by User to `users`
func (*Role) TableName() string {
	return "roles"
}
