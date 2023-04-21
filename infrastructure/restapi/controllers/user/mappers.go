// Package user contains the user controller
package user

import (
	userUseCase "hacktiv/final-project/application/usecases/user"
	domainRole "hacktiv/final-project/domain/role"
	userDomain "hacktiv/final-project/domain/user"
)

func userRoleDomainToResponseMapper(userRoleDomain *domainRole.User) (createUserRoleResponse *ResponseUserRole) {
	createUserRoleResponse = &ResponseUserRole{ID: userRoleDomain.ID, UserName: userRoleDomain.UserName,
		Email: userRoleDomain.Email, FirstName: userRoleDomain.FirstName, LastName: userRoleDomain.LastName, Role: userRoleDomain.Role}

	return
}

func domainToResponseMapper(userDomain *userDomain.User) (createUserResponse *ResponseUser) {
	createUserResponse = &ResponseUser{ID: userDomain.ID, UserName: userDomain.UserName,
		Email: userDomain.Email, FirstName: userDomain.FirstName, LastName: userDomain.LastName,
		CreatedAt: userDomain.CreatedAt, UpdatedAt: userDomain.UpdatedAt}

	return
}

func arrayDomainToResponseMapper(usersDomain *[]userDomain.User) *[]ResponseUser {
	usersResponse := make([]ResponseUser, len(*usersDomain))
	for i, user := range *usersDomain {
		usersResponse[i] = *domainToResponseMapper(&user)
	}
	return &usersResponse
}

func toUsecaseMapper(user *NewUserRequest) userUseCase.NewUser {
	return userUseCase.NewUser{UserName: user.UserName, Password: user.Password, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName, RoleID: user.RoleID}

}
