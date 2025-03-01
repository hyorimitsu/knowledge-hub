package errors

import (
	"net/http"
	"reflect"
	stdErrors "errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator is a custom validator for Echo
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator creates a new custom validator
func NewCustomValidator() *CustomValidator {
	v := validator.New()
	
	// Register custom validation tags here if needed
	
	// Use JSON tag names for validation errors
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	
	return &CustomValidator{
		validator: v,
	}
}

// Validate validates a struct
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Convert validator errors to our custom validation error
		return NewValidationErrorFromValidator(err)
	}
	return nil
}

// BindAndValidate binds and validates a request
func BindAndValidate(c echo.Context, i interface{}) error {
	// Bind request body to struct
	if err := c.Bind(i); err != nil {
		return NewWithMessage(ErrBadRequest, "Invalid request body", err)
	}
	
	// Validate struct
	if err := c.Validate(i); err != nil {
		return err // Already converted to our custom validation error
	}
	
	return nil
}

// RegisterEchoValidator registers the custom validator with Echo
func RegisterEchoValidator(e *echo.Echo) {
	e.Validator = NewCustomValidator()
}

// ValidationErrorHandler is a middleware that handles validation errors
func ValidationErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		
		// Check if it's a validation error
		var validationErr *ValidationError
		if stdErrors.As(err, &validationErr) {
			return RespondWithError(c, validationErr)
		}
		
		// Check if it's an Echo validation error
		if he, ok := err.(*echo.HTTPError); ok && he.Code == http.StatusBadRequest {
			// Convert Echo validation error to our custom error
			return RespondWithError(c, New(ErrBadRequest, err))
		}
		
		// Pass other errors to the next handler
		return err
	}
}