package user

import "time"

// Role is a struct that contains the role information
type Role struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRole is a struct that contains role of user
type UserRole struct {
	User
	Role Role `gorm:"foreignKey:ID;references:RoleID"`
}

// TableName overrides the table name used by User to `users`
func (*UserRole) TableName() string {
	return "users"
}
