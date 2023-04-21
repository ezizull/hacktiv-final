package routes

import (
	userController "hacktiv/final-project/infrastructure/restapi/controllers/user"
	"hacktiv/final-project/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router *gin.RouterGroup, controller *userController.Controller) {
	routerAuth := router.Group("/user")

	// public
	{
		routerAuth.POST("", controller.NewUser)
	}

	// authentication
	routerAuth.Use(middlewares.AuthJWTMiddleware())
	{
		routerAuth.GET("/:id", controller.GetUsersByID)
		routerAuth.PUT("/:id", controller.UpdateUser)
		routerAuth.DELETE("/:id", controller.DeleteUser)
	}

	// admin role
	routerAuth.Use(middlewares.AuthRoleMiddleware([]string{"admin"}))
	{
		routerAuth.GET("", controller.GetAllUsers)
	}
}
