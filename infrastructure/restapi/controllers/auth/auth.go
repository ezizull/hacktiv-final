// Package auth contains the auth controller
package auth

import (
	useCaseAuth "hacktiv/final-project/application/usecases/auth"
	errorDomain "hacktiv/final-project/domain/errors"
	userDomain "hacktiv/final-project/domain/user"
	"hacktiv/final-project/infrastructure/restapi/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the auth service
type Controller struct {
	AuthService useCaseAuth.Service
}

// Login godoc
// @Tags auth
// @Summary Login UserName
// @Description Auth user by email and password
// @Param data body LoginRequest true "body data"
// @Success 200 {object} userDomain.SecurityAuthenticatedUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /auth/login [post]
func (c *Controller) Login(ctx *gin.Context) {
	var request LoginRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	user := userDomain.LoginUser{
		Email:    request.Email,
		Password: request.Password,
	}

	authDataUser, err := c.AuthService.Login(user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// host := strings.Split(ctx.Request.Host, ":")[0] // remove the port if it exists
	// ctx.SetCookie("X-CSRF", "value", int(authDataUser.Security.ExpirationAccessDateTime.Second()), "/", host, false, true)

	ctx.JSON(http.StatusOK, authDataUser.CustomResponse())
}

// GetAccessTokenByRefreshToken godoc
// @Tags auth
// @Summary GetAccessTokenByRefreshToken UserName
// @Description Auth user by email and password
// @Param data body AccessTokenRequest true "body data"
// @Success 200 {object} userDomain.SecurityAuthenticatedUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /auth/access-token [post]
func (c *Controller) GetAccessTokenByRefreshToken(ctx *gin.Context) {
	var request AccessTokenRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	// CSRF, err := ctx.Cookie("X-CSRF")
	// if err != nil {
	// 	appError := errorDomain.NewAppError(err, errorDomain.TokenGeneratorError)
	// 	_ = ctx.Error(appError)
	// 	return
	// }

	authDataUser, err := c.AuthService.AccessTokenByRefreshToken(request.RefreshToken, "CSRF")
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, authDataUser)
}
