package routes

import (
	photoController "hacktiv/final-project/infrastructure/restapi/controllers/photo"
	"hacktiv/final-project/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// PhotoRoutes is a function that contains all routes of the photo
func PhotoRoutes(router *gin.RouterGroup, controller *photoController.Controller) {
	routerPhoto := router.Group("/photos")

	// authentication
	routerPhoto.Use(middlewares.AuthJWTMiddleware())
	{
		routerPhoto.GET("", controller.GetAllPhotos)
		routerPhoto.GET("/own", controller.GetAllOwnPhotos)
		routerPhoto.GET("/:id/comments", controller.GetPhotoWithComments)
		routerPhoto.GET("/:id", controller.GetPhotoByID)
		routerPhoto.POST("", controller.NewPhoto)
		routerPhoto.PUT("/:id", controller.UpdatePhoto)
		routerPhoto.DELETE("/:id", controller.DeletePhoto)
	}
}
