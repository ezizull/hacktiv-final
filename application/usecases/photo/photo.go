package photo

import (
	photoDomain "hacktiv/final-project/domain/photo"

	photoRepository "hacktiv/final-project/infrastructure/repository/postgres/photo"
	userRepository "hacktiv/final-project/infrastructure/repository/postgres/user"
)

// Service is a struct that contains the repository implementation for photo use case
type Service struct {
	PhotoTesting    photoRepository.PhotoTesting
	PhotoRepository photoRepository.Repository
	UserRepository  userRepository.Repository
}

// GetAll is a function that returns all photos
func (s *Service) GetAll(page int64, limit int64) (*photoDomain.PaginationResultPhoto, error) {

	all, err := s.PhotoRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &photoDomain.PaginationResultPhoto{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// UserGetAll is a function that returns all photos
func (s *Service) UserGetAll(page int64, userId int, limit int64) (*photoDomain.PaginationResultPhoto, error) {

	all, err := s.PhotoRepository.UserGetAll(page, userId, limit)
	if err != nil {
		return nil, err
	}

	return &photoDomain.PaginationResultPhoto{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetByID is a function that returns a photo by id
func (s *Service) GetByID(id int) (*photoDomain.Photo, error) {
	return s.PhotoRepository.GetByID(id)
}

// UserGetByID is a function that returns a photo by id
func (s *Service) UserGetByID(id int, userId int) (*photoDomain.Photo, error) {
	return s.PhotoRepository.UserGetByID(id, userId)
}

// Create is a function that creates a photo
func (s *Service) Create(photo *photoDomain.NewPhoto) (*photoDomain.Photo, error) {

	photoModel := photo.ToDomainMapper()

	return s.PhotoRepository.Create(photoModel)
}

// GetByMap is a function that returns a photo by map
func (s *Service) GetByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error) {
	return s.PhotoRepository.GetOneByMap(photoMap)
}

// Delete is a function that deletes a photo by id
func (s *Service) Delete(id int) (err error) {
	return s.PhotoRepository.Delete(id)
}

// Update is a function that updates a photo by id
func (s *Service) Update(id int, updatePhoto photoDomain.UpdatePhoto) (*photoDomain.Photo, error) {
	photo := updatePhoto.ToDomainMapper()
	return s.PhotoRepository.Update(id, &photo)
}
