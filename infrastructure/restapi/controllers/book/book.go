package book

import (
	"errors"
	security "hacktiv/final-project/application/security/jwt"
	useCaseBook "hacktiv/final-project/application/usecases/book"
	bookDomain "hacktiv/final-project/domain/book"
	errorDomain "hacktiv/final-project/domain/errors"
	"hacktiv/final-project/infrastructure/restapi/controllers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	BookService useCaseBook.Service
}

// NewBook godoc
// @Tags book
// @Summary Create New Book
// @Descriptioniption Create new book on the system
// @Accept  json
// @Produce  json
// @Param data body NewBookRequest true "body data"
// @Success 200 {object} bookDomain.Book
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /book [post]
func (c *Controller) NewBook(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Auth").(security.Claims)

	var request NewBookRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	newBook := useCaseBook.NewBook{
		Title:       request.Title,
		UserID:      authData.UserID,
		Description: request.Description,
	}

	bookDomain, err := c.BookService.Create(&newBook)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, bookDomain)
}

// GetAllBooks godoc
// @Tags book
// @Summary Get all Books
// @Description Get all Books on the system
// @Param   limit  query   string  true        "limit"
// @Param   page  query   string  true        "page"
// @Success 200 {object} []bookDomain.PaginationResultBook
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /book [get]
func (c *Controller) GetAllBooks(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Auth").(security.Claims)

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "20")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param page is necessary to be an integer"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param limit is necessary to be an integer"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var books *bookDomain.PaginationResultBook

	if authData.Role == "admin" {
		books, err = c.BookService.GetAll(page, limit)
		if err != nil {
			appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			_ = ctx.Error(appError)
			return
		}
	} else {
		books, err = c.BookService.UserGetAll(page, authData.UserID, limit)
		if err != nil {
			appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			_ = ctx.Error(appError)
			return
		}
	}

	ctx.JSON(http.StatusOK, books)
}

// GetBookByID godoc
// @Tags book
// @Summary Get books by ID
// @Descriptioniption Get Books by ID on the system
// @Param book_id path int true "id of book"
// @Success 200 {object} bookDomain.Book
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /book/{book_id} [get]
func (c *Controller) GetBookByID(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Auth").(security.Claims)

	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("book id is invalid"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var bookDomain *bookDomain.Book

	if authData.Role == "admin" {
		bookDomain, err = c.BookService.GetByID(bookID)
		if err != nil {
			appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
			_ = ctx.Error(appError)
			return
		}
	} else {

		bookDomain, err = c.BookService.UserGetByID(bookID, authData.UserID)
		if err != nil {
			appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
			_ = ctx.Error(appError)
			return
		}
	}

	ctx.JSON(http.StatusOK, bookDomain)
}

// UpdateBook is the controller to update a book
func (c *Controller) UpdateBook(ctx *gin.Context) {
	bookID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	var requestMap map[string]interface{}

	err = controllers.BindJSONMap(ctx, &requestMap)
	if err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = updateValidation(requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var book *bookDomain.Book
	book, err = c.BookService.Update(uint(bookID), requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, book)

}

// DeleteBook is the controller to delete a book
func (c *Controller) DeleteBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.BookService.Delete(bookID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
