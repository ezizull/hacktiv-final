// Package routes contains all routes of the application
package routes

import (
	// swaggerFiles for documentation
	_ "hacktiv/final-project/docs"
	"hacktiv/final-project/infrastructure/restapi/adapter"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// Security is a struct that contains the security of the application
// @SecurityDefinitions.jwt
type Security struct {
	Authorization string `header:"Authorization" json:"Authorization"`
}

func ApplicationRootRouter(router *gin.Engine, db *gorm.DB) {
	// Documentation Swagger
	{
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
	}
}

func ApplicationV1Router(router *gin.Engine, db *gorm.DB) {
	routerV1 := router.Group("/v1")

	{
		// Auth Routes
		AuthRoutes(routerV1, adapter.AuthAdapter(db))

		// User Routes
		UserRoutes(routerV1, adapter.UserAdapter(db))

		// Photo Routes
		PhotoRoutes(routerV1, adapter.PhotoAdapter(db))

		// SocialMedia Routes
		SocialMediaRoutes(routerV1, adapter.SocialMediaAdapter(db))

		// Comment Routes
		CommentRoutes(routerV1, adapter.CommentAdapter(db))

	}
}
