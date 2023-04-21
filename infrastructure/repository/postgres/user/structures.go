// Package user contains the business logic for the user entity
package user

import (
	"hacktiv/final-project/infrastructure/repository/postgres/role"
	"time"
)

// User is a struct that contains the user information
type User struct {
	ID           int       `json:"id" example:"1099" gorm:"primaryKey"`
	UserName     string    `json:"userName" example:"UserName" gorm:"column:user_name;unique"`
	Email        string    `json:"email" example:"some@mail.com" gorm:"unique"`
	FirstName    string    `json:"first_name" example:"John"`
	LastName     string    `json:"last_name" example:"Doe"`
	HashPassword string    `json:"hash_password" example:"SomeHashPass"`
	RoleID       string    `json:"role_id" gorm:"index"`
	CreatedAt    time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

// TableName overrides the table name used by User to `users`
func (*User) TableName() string {
	return "users"
}

// PaginationResultUser is a struct that contains the pagination result for user
type PaginationResultUser struct {
	Data       []User
	Total      int
	Limit      int
	Current    int
	NextCursor uint
	PrevCursor uint
	NumPages   int
}

// UserRole is a struct that contains role of user
type UserRole struct {
	User
	Role role.Role `gorm:"foreignKey:ID;references:RoleID"`
}

// TableName overrides the table name used by User to `users`
func (*UserRole) TableName() string {
	return "users"
}
