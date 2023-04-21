package role

import (
	domainError "hacktiv/final-project/domain/errors"
	domainRole "hacktiv/final-project/domain/role"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for user entity
type Repository struct {
	DB *gorm.DB
}

// GetByID ... Fetch only one role by ID
func (r *Repository) GetByID(id string) (*domainRole.Role, error) {
	var role Role
	err := r.DB.Where("id = ?", id).First(&role).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainError.NewAppErrorWithType(domainError.NotFound)
		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
	}

	return role.toDomainMapper(), err
}
