// Package user provides the use case for user
package user

import (
	domainUser "hacktiv/final-project/domain/user"
)

func (n *NewUser) toDomainMapper() *domainUser.User {
	return &domainUser.User{
		UserName:  n.UserName,
		Email:     n.Email,
		FirstName: n.FirstName,
		LastName:  n.LastName,
		RoleID:    n.RoleID,
	}
}
