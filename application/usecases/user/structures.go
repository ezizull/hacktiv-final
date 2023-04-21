// Package user provides the use case for user
package user

import (
	domainUser "hacktiv/final-project/domain/user"
)

// NewUser is the structure for a new user
type NewUser struct {
	UserName  string
	Email     string
	FirstName string
	LastName  string
	Password  string
	RoleID    string
}

// PaginationResultUser is the structure for pagination result of user
type PaginationResultUser struct {
	Data       []domainUser.User
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
