package errors

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationErrorsToMap converts validator.ValidationErrors to a map of field:message
func ValidationErrorsToMap(err error) map[string]string {
	fieldErrors := make(map[string]string)
	
	if err == nil {
		return fieldErrors
	}
	
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		// If it's not a validator.ValidationErrors, just return a generic error
		fieldErrors["general"] = err.Error()
		return fieldErrors
	}
	
	for _, e := range validationErrors {
		// Convert the field name to JSON format (usually lowercase first letter)
		field := e.Field()
		if len(field) > 0 {
			field = strings.ToLower(field[:1]) + field[1:]
		}
		
		// Create a user-friendly error message based on the validation tag
		var message string
		switch e.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", field)
		case "email":
			message = fmt.Sprintf("%s must be a valid email address", field)
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters long", field, e.Param())
		case "max":
			message = fmt.Sprintf("%s must be at most %s characters long", field, e.Param())
		case "oneof":
			message = fmt.Sprintf("%s must be one of: %s", field, e.Param())
		case "uuid":
			message = fmt.Sprintf("%s must be a valid UUID", field)
		case "url":
			message = fmt.Sprintf("%s must be a valid URL", field)
		default:
			// For other validation tags, use a generic message
			message = fmt.Sprintf("%s failed validation: %s", field, e.Tag())
		}
		
		fieldErrors[field] = message
	}
	
	return fieldErrors
}

// NewValidationErrorFromValidator creates a ValidationError from validator.ValidationErrors
func NewValidationErrorFromValidator(err error) *ValidationError {
	fieldErrors := ValidationErrorsToMap(err)
	return NewValidationError("Validation failed", fieldErrors, err)
}