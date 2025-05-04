package errors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ValidationError(field, message string) string {
	return fmt.Sprintf("Field %s failed validation: %s", field, message)
}

func GetValidationErrors(errs validator.ValidationErrors) []string {
	var errors []string
	for _, err := range errs {
		errors = append(errors, ValidationError(err.Field(), err.Tag()))
	}

	return errors
}
