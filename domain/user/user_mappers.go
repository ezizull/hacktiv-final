// Package user contains the business logic for the user entity
package user

// ToDomainMapper function to convert user repo to user domain
func (user *User) ToDomainMapper() *User {
	return &User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		HashPassword: user.HashPassword,
		RoleID:       user.RoleID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

// ArrayToDomainMapper function to convert list user domain to list user repo
func ArrayToDomainMapper(users *[]User) *[]User {
	usersDomain := make([]User, len(*users))
	for i, user := range *users {
		usersDomain[i] = *user.ToDomainMapper()
	}

	return &usersDomain
}
