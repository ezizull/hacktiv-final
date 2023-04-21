// Package user contains the business logic for the user entity
package user

import (
	"encoding/json"

	domainError "hacktiv/final-project/domain/errors"
	domainRole "hacktiv/final-project/domain/role"
	domainUser "hacktiv/final-project/domain/user"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for user entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all user data
func (r *Repository) GetAll() (*[]domainUser.User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	if err != nil {
		err = domainError.NewAppErrorWithType(domainError.UnknownError)
		return nil, err
	}

	return arrayToDomainMapper(&users), err
}

// Create ... Insert New data
func (r *Repository) Create(userDomain *domainUser.User) (*domainUser.User, error) {
	userRepository := fromDomainMapper(userDomain)
	txDb := r.DB.Create(userRepository)
	err := txDb.Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError domainError.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &domainUser.User{}, err
		}
		switch newError.Number {
		case 1062:
			err = domainError.NewAppErrorWithType(domainError.ResourceAlreadyExists)
			return &domainUser.User{}, err

		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
	}
	return userRepository.toDomainMapper(), err
}

// GetOneByMap ... Fetch only one user by Map values
func (r *Repository) GetOneByMap(userMap map[string]interface{}) (*domainUser.User, error) {
	var userRepository User

	tx := r.DB.Where(userMap).Limit(1).Find(&userRepository)
	if tx.Error != nil {
		err := domainError.NewAppErrorWithType(domainError.UnknownError)
		return &domainUser.User{}, err
	}
	return userRepository.toDomainMapper(), nil
}

// GetWithRoleByMap ... Fetch only one user with Role by Map values
func (r *Repository) GetWithRoleByMap(userMap map[string]interface{}) (*domainRole.User, error) {
	var userRole UserRole

	err := r.DB.Preload("Role").Where(userMap).First(&userRole).Error
	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainError.NewAppErrorWithType(domainError.NotFound)
		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
	}

	return userRole.toDomainMapper(), nil
}

// GetWithRole ... Fetch only one user with Role by ID
func (r *Repository) GetWithRole(id int) (*domainRole.User, error) {
	var userRole UserRole
	err := r.DB.Preload("Role").Where("id = ?", id).First(&userRole).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainError.NewAppErrorWithType(domainError.NotFound)
		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
	}

	return userRole.toDomainMapper(), err
}

// GetByID ... Fetch only one user by ID
func (r *Repository) GetByID(id int) (*domainUser.User, error) {
	var user User
	err := r.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainError.NewAppErrorWithType(domainError.NotFound)
		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
	}

	return user.toDomainMapper(), err
}

// Update ... Update user
func (r *Repository) Update(id int, userMap map[string]interface{}) (*domainUser.User, error) {
	var user User

	user.ID = id
	err := r.DB.Model(&user).
		Select("user", "email", "firstName", "lastName").
		Updates(userMap).Error

	// err = config.DB.Save(user).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError domainError.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &domainUser.User{}, err
		}
		switch newError.Number {
		case 1062:
			err = domainError.NewAppErrorWithType(domainError.ResourceAlreadyExists)
		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
		return &domainUser.User{}, err

	}

	err = r.DB.Where("id = ?", id).First(&user).Error

	return user.toDomainMapper(), err
}

// Delete ... Delete user
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&User{}, id)
	if tx.Error != nil {
		err = domainError.NewAppErrorWithType(domainError.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = domainError.NewAppErrorWithType(domainError.NotFound)
	}

	return
}
