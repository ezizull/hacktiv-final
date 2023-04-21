package role

import "time"

// User is a struct that contains the user information
type User struct {
	ID           int
	UserName     string
	Email        string
	HashPassword string
	FirstName    string
	LastName     string
	RoleID       string
	Role         Role `gorm:"foreignKey:ID;references:RoleID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Role is a struct that contains the role information
type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
