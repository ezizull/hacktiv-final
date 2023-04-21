package photo

import (
	photoDomain "hacktiv/final-project/domain/photo"
	photoRepository "hacktiv/final-project/infrastructure/repository/postgres/photo"
)

type PhotoTesting interface {
	GetAll(page int64, limit int64) (*photoDomain.PaginationResultPhoto, error)
	UserGetAll(page int64, userId int, limit int64) (*photoDomain.PaginationResultPhoto, error)
	GetByID(id int) (*photoDomain.Photo, error)
	UserGetByID(id int, userId int) (*photoDomain.Photo, error)
	Create(photo *photoDomain.NewPhoto) (*photoDomain.Photo, error)
	GetByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error)
	Delete(id int) (err error)
	Update(id int, updatePhoto photoDomain.UpdatePhoto) (*photoDomain.Photo, error)
}

func NewTesting(photoTest photoRepository.PhotoTesting) PhotoTesting {
	return &Service{
		PhotoTesting: photoTest,
	}
}
