package comment

import (
	commentDomain "hacktiv/final-project/domain/comment"

	commentRepository "hacktiv/final-project/infrastructure/repository/postgres/comment"
)

// Service is a struct that contains the repository implementation for comment use case
type Service struct {
	CommentTesting    commentRepository.CommentTesting
	CommentRepository commentRepository.Repository
}

// GetAll is a function that returns all comments
func (s *Service) GetAll(page int64, limit int64) (*commentDomain.PaginationResultComment, error) {

	all, err := s.CommentRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &commentDomain.PaginationResultComment{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// UserGetAll is a function that returns all comments
func (s *Service) UserGetAll(page int64, userId int, limit int64) (*commentDomain.PaginationResultComment, error) {

	all, err := s.CommentRepository.UserGetAll(page, userId, limit)
	if err != nil {
		return nil, err
	}

	return &commentDomain.PaginationResultComment{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetByID is a function that returns a comment by id
func (s *Service) GetByID(id int) (*commentDomain.Comment, error) {
	return s.CommentRepository.GetByID(id)
}

// UserGetByID is a function that returns a comment by id
func (s *Service) UserGetByID(id int, userId int) (*commentDomain.Comment, error) {
	return s.CommentRepository.UserGetByID(id, userId)
}

// Create is a function that creates a comment
func (s *Service) Create(comment *commentDomain.NewComment) (*commentDomain.Comment, error) {

	commentModel := comment.ToDomainMapper()

	return s.CommentRepository.Create(commentModel)
}

// GetByMap is a function that returns a comment by map
func (s *Service) GetByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error) {
	return s.CommentRepository.GetOneByMap(commentMap)
}

// Delete is a function that deletes a comment by id
func (s *Service) Delete(id int) (err error) {
	return s.CommentRepository.Delete(id)
}

// Update is a function that updates a comment by id
func (s *Service) Update(id int, updateComment commentDomain.UpdateComment) (*commentDomain.Comment, error) {
	comment := updateComment.ToDomainMapper()
	return s.CommentRepository.Update(id, &comment)
}

// Update is a function that updates a comment by id
func (s *Service) UserUpdate(id int, userId int, updateComment commentDomain.UpdateComment) (*commentDomain.Comment, error) {
	comment := updateComment.ToDomainMapper()
	return s.CommentRepository.UserUpdate(id, userId, &comment)
}
