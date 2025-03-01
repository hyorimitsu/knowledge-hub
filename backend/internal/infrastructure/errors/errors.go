package errors

import (
	"fmt"
	"strings"
)

// AppError is the base error type for application errors
type AppError struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
	Err     error       `json:"-"` // Original error (not serialized)
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// StatusCode returns the HTTP status code for the error
func (e *AppError) StatusCode() int {
	return e.Code.HTTPStatusCode()
}

// New creates a new AppError with the given code and optional wrapped error
func New(code ErrorCode, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: code.DefaultMessage(),
		Err:     err,
	}
}

// NewWithMessage creates a new AppError with the given code, message and optional wrapped error
func NewWithMessage(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// NewWithDetails creates a new AppError with the given code, message, details and optional wrapped error
func NewWithDetails(code ErrorCode, message string, details interface{}, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
		Err:     err,
	}
}

// ValidationError represents a validation error with field-specific details
type ValidationError struct {
	*AppError
	FieldErrors map[string]string `json:"fieldErrors,omitempty"`
}

// NewValidationError creates a new ValidationError
func NewValidationError(message string, fieldErrors map[string]string, err error) *ValidationError {
	if message == "" {
		message = "Validation failed"
	}
	
	return &ValidationError{
		AppError: &AppError{
			Code:    ErrValidation,
			Message: message,
			Err:     err,
		},
		FieldErrors: fieldErrors,
	}
}

// NewValidationErrorFromStrings creates a ValidationError from a slice of error strings
func NewValidationErrorFromStrings(errors []string, err error) *ValidationError {
	return &ValidationError{
		AppError: &AppError{
			Code:    ErrValidation,
			Message: "Validation failed",
			Err:     err,
		},
		FieldErrors: parseValidationErrors(errors),
	}
}

// parseValidationErrors converts string errors like "field: error message" to a map
func parseValidationErrors(errors []string) map[string]string {
	fieldErrors := make(map[string]string)
	
	for _, err := range errors {
		parts := strings.SplitN(err, ":", 2)
		if len(parts) == 2 {
			field := strings.TrimSpace(parts[0])
			message := strings.TrimSpace(parts[1])
			fieldErrors[field] = message
		} else {
			// If the error doesn't follow the field:message format, use it as is
			fieldErrors["general"] = err
		}
	}
	
	return fieldErrors
}

// DomainError represents an error in the domain layer
type DomainError struct {
	*AppError
}

// NewDomainError creates a new DomainError
func NewDomainError(code ErrorCode, message string, err error) *DomainError {
	if message == "" {
		message = code.DefaultMessage()
	}
	
	return &DomainError{
		AppError: &AppError{
			Code:    code,
			Message: message,
			Err:     err,
		},
	}
}

// ApplicationError represents an error in the application layer
type ApplicationError struct {
	*AppError
}

// NewApplicationError creates a new ApplicationError
func NewApplicationError(code ErrorCode, message string, err error) *ApplicationError {
	if message == "" {
		message = code.DefaultMessage()
	}
	
	return &ApplicationError{
		AppError: &AppError{
			Code:    code,
			Message: message,
			Err:     err,
		},
	}
}

// InfrastructureError represents an error in the infrastructure layer
type InfrastructureError struct {
	*AppError
}

// NewInfrastructureError creates a new InfrastructureError
func NewInfrastructureError(code ErrorCode, message string, err error) *InfrastructureError {
	if message == "" {
		message = code.DefaultMessage()
	}
	
	return &InfrastructureError{
		AppError: &AppError{
			Code:    code,
			Message: message,
			Err:     err,
		},
	}
}