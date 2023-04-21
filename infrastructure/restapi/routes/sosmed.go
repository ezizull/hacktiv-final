package routes

import (
	sosmedController "hacktiv/final-project/infrastructure/restapi/controllers/sosmed"
	"hacktiv/final-project/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// SocialMediaRoutes is a function that contains all routes of the sosmed
func SocialMediaRoutes(router *gin.RouterGroup, controller *sosmedController.Controller) {
	routerSocialMedia := router.Group("/social-media")

	// authentication
	routerSocialMedia.Use(middlewares.AuthJWTMiddleware())
	{
		routerSocialMedia.GET("", controller.GetAllSocialMedia)
		routerSocialMedia.GET("/:id", controller.GetSocialMediaByID)
		routerSocialMedia.POST("", controller.NewSocialMedia)

	}

	// admin role
	routerSocialMedia.Use(middlewares.AuthRoleMiddleware([]string{"admin"}))
	{
		routerSocialMedia.PUT("/:id", controller.UpdateSocialMedia)
		routerSocialMedia.DELETE("/:id", controller.DeleteSocialMedia)
	}

}
