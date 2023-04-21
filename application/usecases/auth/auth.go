// Package auth provides the use case for authentication
package auth

import (
	"errors"
	"fmt"
	"time"

	"hacktiv/final-project/application/security/jwt"
	errorDomain "hacktiv/final-project/domain/errors"
	userDomain "hacktiv/final-project/domain/user"

	userRepository "hacktiv/final-project/infrastructure/repository/postgres/user"

	"golang.org/x/crypto/bcrypt"
)

// Service is a struct that contains the repository implementation for auth use case
type Service struct {
	UserRepository userRepository.Repository
}

// Login implements the login use case
func (s *Service) Login(user userDomain.LoginUser) (*userDomain.SecurityAuthenticatedUser, error) {
	userMap := map[string]interface{}{"email": user.Email}
	userRole, err := s.UserRepository.GetWithRoleByMap(userMap)
	fmt.Println("check ", userRole)
	if err != nil {
		return &userDomain.SecurityAuthenticatedUser{}, err
	}
	if userRole.ID == 0 {
		return &userDomain.SecurityAuthenticatedUser{}, errorDomain.NewAppError(errors.New("email or password does not match"), errorDomain.NotAuthorized)
	}

	isAuthenticated := CheckPasswordHash(user.Password, userRole.HashPassword)
	if !isAuthenticated {
		err = errorDomain.NewAppError(err, errorDomain.NotAuthorized)
		return &userDomain.SecurityAuthenticatedUser{}, errorDomain.NewAppError(errors.New("email or password does not match"), errorDomain.NotAuthorized)
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(userRole.ID, "access", userRole.Role.Name)
	if err != nil {
		return &userDomain.SecurityAuthenticatedUser{}, err
	}
	refreshTokenClaims, err := jwt.GenerateJWTToken(userRole.ID, "refresh", userRole.Role.Name)
	if err != nil {
		return &userDomain.SecurityAuthenticatedUser{}, err
	}

	return userDomain.SecAuthUserRoleMapper(userRole, &userDomain.Auth{
		AccessToken:               accessTokenClaims.Token,
		RefreshToken:              refreshTokenClaims.Token,
		ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
		ExpirationRefreshDateTime: refreshTokenClaims.ExpirationTime,
	}), err
}

// AccessTokenByRefreshToken implements the Access Token By Refresh Token use case
func (s *Service) AccessTokenByRefreshToken(refreshToken string) (*userDomain.SecurityAuthenticatedUser, error) {
	claimsMap, err := jwt.GetClaimsAndVerifyToken(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	userMap := map[string]interface{}{"id": claimsMap["user_id"]}
	userRole, err := s.UserRepository.GetWithRoleByMap(userMap)
	if err != nil {
		return nil, err

	}
	if userRole.ID == 0 {
		return &userDomain.SecurityAuthenticatedUser{}, errorDomain.NewAppError(errors.New(errorDomain.TokenGeneratorErrorMessage), errorDomain.NotFound)
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(userRole.ID, "access", userRole.Role.Name)
	if err != nil {
		return &userDomain.SecurityAuthenticatedUser{}, err
	}

	var expTime = int64(claimsMap["exp"].(float64))

	return userDomain.SecAuthUserRoleMapper(userRole, &userDomain.Auth{
		AccessToken:               accessTokenClaims.Token,
		ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
		RefreshToken:              refreshToken,
		ExpirationRefreshDateTime: time.Unix(expTime, 0),
	}), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
