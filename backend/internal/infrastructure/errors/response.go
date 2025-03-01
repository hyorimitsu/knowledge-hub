package errors

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

// ErrorResponse is the standardized error response format
type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   *ErrorData  `json:"error"`
}

// ErrorData contains detailed error information
type ErrorData struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// RespondWithError sends a standardized error response
func RespondWithError(c echo.Context, err error) error {
	// Default to internal server error
	statusCode := http.StatusInternalServerError
	errorCode := ErrInternal
	message := "An internal server error occurred"
	var details interface{}

	// Try to get more specific error information
		var validationErr *ValidationError
		if errors.As(err, &validationErr) {
			statusCode = validationErr.StatusCode()
			errorCode = validationErr.Code
			message = validationErr.Message
			details = validationErr.FieldErrors
		} else {
			var appErr *AppError
			if errors.As(err, &appErr) {
				statusCode = appErr.StatusCode()
				errorCode = appErr.Code
				message = appErr.Message
				details = appErr.Details
			}
		}

	// Log the error
	logError(c, err, statusCode)

	// Create the response
	response := ErrorResponse{
		Success: false,
		Error: &ErrorData{
			Code:    string(errorCode),
			Message: message,
			Details: details,
		},
	}

	return c.JSON(statusCode, response)
}

// logError logs the error with appropriate level based on status code
func logError(c echo.Context, err error, statusCode int) {
	// Get request information
	req := c.Request()
	path := req.URL.Path
	method := req.Method

	// Format the log message
	logMsg := fmt.Sprintf("[%s] %s - Error: %v", method, path, err)

	// Log with appropriate level based on status code
	logger := c.Logger()

	if statusCode >= 500 {
		// For server errors, log at error level with stack trace
		logger.Error(logMsg)
		logger.Error(string(debug.Stack()))
	} else if statusCode >= 400 {
		// For client errors, log at warn level
		logger.Warn(logMsg)
	} else {
		// For other status codes, log at info level
		logger.Info(logMsg)
	}
}

// SuccessResponse is the standardized success response format
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

// RespondWithSuccess sends a standardized success response
func RespondWithSuccess(c echo.Context, statusCode int, data interface{}) error {
	response := SuccessResponse{
		Success: true,
		Data:    data,
	}

	return c.JSON(statusCode, response)
}

// RespondWithEmpty sends a success response with no data
func RespondWithEmpty(c echo.Context, statusCode int) error {
	return c.JSON(statusCode, SuccessResponse{
		Success: true,
	})
}