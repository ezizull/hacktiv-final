// Package user contains the user controller
package user

import (
	"errors"
	"fmt"
	errorDomain "hacktiv/final-project/domain/errors"
	userDomain "hacktiv/final-project/domain/user"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func updateValidation(request map[string]interface{}) (err error) {
	var errorsValidation []string

	for k, v := range request {
		if v == "" {
			errorsValidation = append(errorsValidation, fmt.Sprintf("%s cannot be empty", k))
		}
	}

	if password, ok := request["password"].(string); ok {
		if len(password) < 8 {
			errorsValidation = append(errorsValidation, "password must be at least 8 characters long")
		}
		hasSpecialChar := regexp.MustCompile(`[^a-zA-Z0-9]+`).MatchString
		if !hasSpecialChar(password) {
			errorsValidation = append(errorsValidation, "password must contain at least one special character")
		}
		hasCapitalLetter := regexp.MustCompile(`[A-Z]+`).MatchString
		if !hasCapitalLetter(password) {
			errorsValidation = append(errorsValidation, "password must contain at least one capital letter")
		}
		hasLowerCase := regexp.MustCompile(`[a-z]+`).MatchString
		if !hasLowerCase(password) {
			errorsValidation = append(errorsValidation, "password must contain at least one lowercase letter")
		}
		hasNumber := regexp.MustCompile(`[0-9]+`).MatchString
		if !hasNumber(password) {
			errorsValidation = append(errorsValidation, "password must contain at least one number")
		}
	}

	validationMap := map[string]string{
		"username":  "omitempty,gt=3,lt=100",
		"email":     "omitempty,gt=3,lt=100,email",
		"firstName": "omitempty,gt=2,lt=100",
		"lastName":  "omitempty,gt=2,lt=100",
		"role_id":   "omitempty,gt=0,lt=100",
	}

	validate := validator.New()
	err = validate.RegisterValidation("update_validation", func(fl validator.FieldLevel) bool {
		m, ok := fl.Field().Interface().(map[string]interface{})
		if !ok {
			return false
		}

		for k, v := range validationMap {
			errValidate := validate.Var(m[k], v)
			if errValidate != nil {
				validatorErr := errValidate.(validator.ValidationErrors)
				errorsValidation = append(errorsValidation, fmt.Sprintf("%s do not satisfy condition %v=%v", k, validatorErr[0].Tag(), validatorErr[0].Param()))
			}
		}

		return true
	})

	if err != nil {
		err = errorDomain.NewAppError(err, errorDomain.UnknownError)
		return
	}

	err = validate.Var(request, "update_validation")
	if err != nil {
		err = errorDomain.NewAppError(err, errorDomain.UnknownError)
		return
	}
	if errorsValidation != nil {
		err = errorDomain.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), errorDomain.ValidationError)
	}
	return
}

func createValidation(request userDomain.NewUser) error {
	// Username must have minimum length of 4
	if len(request.UserName) < 4 {
		return errors.New("Username must be at least 4 characters long")
	}

	// Email must be a valid email format
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(request.Email) {
		return errors.New("Invalid email format")
	}

	// Password must have minimum length of 8, at least 1 special character, 1 capital letter, 1 lowercase letter, and 1 number
	if len(request.Password) < 8 {
		return errors.New("Password should be at least 8 characters long")
	}
	if !regexp.MustCompile(`[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]`).MatchString(request.Password) {
		return errors.New("Password should contain at least one special character")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(request.Password) {
		return errors.New("Password should contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(request.Password) {
		return errors.New("Password should contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`\d`).MatchString(request.Password) {
		return errors.New("Password should contain at least one number")
	}

	return nil
}
