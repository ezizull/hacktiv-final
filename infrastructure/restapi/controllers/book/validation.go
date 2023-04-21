// Package book contains the book controller
package book

import (
	"errors"
	"fmt"
	"strings"

	errorDomain "hacktiv/final-project/domain/errors"

	"github.com/go-playground/validator/v10"
)

func updateValidation(request map[string]interface{}) (err error) {
	var errorsValidation []string

	for k, v := range request {
		if v == "" {
			errorsValidation = append(errorsValidation, fmt.Sprintf("%s cannot be empty", k))
		}
	}

	validationMap := map[string]string{
		"title":       "omitempty,gt=3,lt=100",
		"description": "omitempty,gt=3,lt=100",
		"author":      "omitempty,gt=3,lt=100",
	}

	validate := validator.New()
	err = validate.RegisterValidation("", func(fl validator.FieldLevel) bool {
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