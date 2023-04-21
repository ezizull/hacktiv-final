package routes

import (
	commentController "hacktiv/final-project/infrastructure/restapi/controllers/comment"
	"hacktiv/final-project/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// CommentRoutes is a function that contains all routes of the comment
func CommentRoutes(router *gin.RouterGroup, controller *commentController.Controller) {
	routerComment := router.Group("/comments")

	// authentication
	routerComment.Use(middlewares.AuthJWTMiddleware())
	{
		routerComment.GET("", controller.GetAllComments)
		routerComment.GET("/own", controller.GetAllOwnComments)
		routerComment.GET("/:id", controller.GetCommentByID)
		routerComment.POST("", controller.NewComment)
		routerComment.PUT("/:id", controller.UpdateComment)
		routerComment.DELETE("/:id", controller.DeleteComment)
	}
}
