package lib

import (
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	Validator  *validator.Validate
	Translator *ut.Translator
}

func (cv *RequestValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		// Create an array of errors
		errors := make([]string, len(validationErrors))

		// Loop through the errors and add the translated error message to the array
		for i, e := range validationErrors {
			errors[i] = e.Translate(*cv.Translator)
		}

		// Return the errors
		return CreateNewResponseError(http.StatusBadRequest, errors...)
	}
	return nil
}
