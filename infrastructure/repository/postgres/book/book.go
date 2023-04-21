package book

import (
	"encoding/json"
	bookDomain "hacktiv/final-project/domain/book"
	errorDomain "hacktiv/final-project/domain/errors"
	"log"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for photo entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all book data
func (r *Repository) GetAll(page int64, limit int64) (*bookDomain.PaginationResultBook, error) {
	var books []Book
	var total int64

	err := r.DB.Model(&Book{}).Count(&total).Error
	if err != nil {
		return &bookDomain.PaginationResultBook{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&books).Error

	if err != nil {
		return &bookDomain.PaginationResultBook{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &bookDomain.PaginationResultBook{
		Data:       arrayToDomainMapper(&books),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// UserGetAll Fetch all book data
func (r *Repository) UserGetAll(page int64, userId int, limit int64) (*bookDomain.PaginationResultBook, error) {
	var books []Book
	var total int64

	err := r.DB.Model(&Book{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return &bookDomain.PaginationResultBook{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&books).Error

	if err != nil {
		return &bookDomain.PaginationResultBook{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &bookDomain.PaginationResultBook{
		Data:       arrayToDomainMapper(&books),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// Create ... Insert New data
func (r *Repository) Create(newBook *bookDomain.Book) (createdBook *bookDomain.Book, err error) {
	book := fromDomainMapper(newBook)

	tx := r.DB.Create(book)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return
	}

	createdBook = book.toDomainMapper()
	return
}

// GetByID ... Fetch only one book by Id
func (r *Repository) GetByID(id int) (*bookDomain.Book, error) {
	var book Book
	err := r.DB.Where("id = ?", id).First(&book).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &bookDomain.Book{}, err
	}

	return book.toDomainMapper(), nil
}

// UserGetByID ... Fetch only one book by Id
func (r *Repository) UserGetByID(id int, userId int) (*bookDomain.Book, error) {
	var book Book
	err := r.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&book).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &bookDomain.Book{}, err
	}

	return book.toDomainMapper(), nil
}

// GetOneByMap ... Fetch only one book by Map
func (r *Repository) GetOneByMap(bookMap map[string]interface{}) (*bookDomain.Book, error) {
	var book Book

	err := r.DB.Where(bookMap).Limit(1).Find(&book).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return nil, err
	}
	return book.toDomainMapper(), err
}

// Update ... Update book
func (r *Repository) Update(id uint, bookMap map[string]interface{}) (*bookDomain.Book, error) {
	var book Book

	book.ID = id
	err := r.DB.Model(&book).
		Select("title", "author", "description").
		Updates(bookMap).Error

	// err = config.DB.Save(book).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &bookDomain.Book{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
			return &bookDomain.Book{}, err

		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			return &bookDomain.Book{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&book).Error

	return book.toDomainMapper(), err
}

// Delete ... Delete book
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&bookDomain.Book{}, id)

	log.Println("check ", tx)
	if tx.Error != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
	}

	return
}
