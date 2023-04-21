package adapter

import (
	commentService "hacktiv/final-project/application/usecases/comment"
	commentRepository "hacktiv/final-project/infrastructure/repository/postgres/comment"
	commentController "hacktiv/final-project/infrastructure/restapi/controllers/comment"

	"gorm.io/gorm"
)

// CommentAdapter is a function that returns a comment controller
func CommentAdapter(db *gorm.DB) *commentController.Controller {
	mRepository := commentRepository.Repository{DB: db}
	service := commentService.Service{CommentRepository: mRepository}
	return &commentController.Controller{CommentService: service}
}
