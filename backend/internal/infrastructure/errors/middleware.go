package errors

import (
	"fmt"
	"net/http"
	"runtime/debug"
	stdErrors "errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ErrorHandlerConfig contains configuration for the error handler middleware
type ErrorHandlerConfig struct {
	// Skipper defines a function to skip middleware
	Skipper middleware.Skipper

	// LogLevel sets the logging level for errors
	LogLevel uint
}

// DefaultErrorHandlerConfig is the default error handler middleware config
var DefaultErrorHandlerConfig = ErrorHandlerConfig{
	Skipper: middleware.DefaultSkipper,
}

// ErrorHandler returns a middleware that handles panics and errors from handlers
func ErrorHandler() echo.MiddlewareFunc {
	return ErrorHandlerWithConfig(DefaultErrorHandlerConfig)
}

// ErrorHandlerWithConfig returns a middleware with config that handles panics and errors from handlers
func ErrorHandlerWithConfig(config ErrorHandlerConfig) echo.MiddlewareFunc {
	// Use default config if provided config is empty
	if config.Skipper == nil {
		config.Skipper = DefaultErrorHandlerConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip if needed
			if config.Skipper(c) {
				return next(c)
			}

			// Execute the next handler
			err := next(c)
			if err != nil {
				// Handle the error
				c.Echo().Logger.Error(err)
				return RespondWithError(c, err)
			}

			return nil
		}
	}
}

// RecoverWithConfig returns a middleware with config that recovers from panics
func RecoverWithConfig() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					// Log the stack trace
					stack := debug.Stack()
					c.Logger().Error(fmt.Sprintf("[PANIC RECOVER] %v\n%s", err, stack))

					// Create an internal server error
					appErr := New(ErrInternal, err)

					// Respond with error
					_ = RespondWithError(c, appErr)
				}
			}()
			return next(c)
		}
	}
}

// ValidationMiddleware returns a middleware that validates request bodies
func ValidationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// The actual validation will be done in the handlers
			// This middleware is a placeholder for future validation logic
			return next(c)
		}
	}
}

// NotFoundHandler handles 404 errors
func NotFoundHandler(c echo.Context) error {
	return RespondWithError(c, New(ErrNotFound, nil))
}

// MethodNotAllowedHandler handles 405 errors
func MethodNotAllowedHandler(c echo.Context) error {
	return RespondWithError(c, NewWithMessage(ErrBadRequest, "Method not allowed", nil))
}

// RegisterErrorHandlers registers custom error handlers with Echo
func RegisterErrorHandlers(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// Handle Echo's built-in errors
		var appErr *AppError
		if !stdErrors.As(err, &appErr) {
			he, ok := err.(*echo.HTTPError)
			if ok {
				// Convert Echo HTTP errors to our AppError format
				switch he.Code {
				case http.StatusNotFound:
					err = New(ErrNotFound, err)
				case http.StatusMethodNotAllowed:
					err = NewWithMessage(ErrBadRequest, "Method not allowed", err)
				case http.StatusBadRequest:
					err = New(ErrBadRequest, err)
				case http.StatusUnauthorized:
					err = New(ErrUnauthorized, err)
				case http.StatusForbidden:
					err = New(ErrForbidden, err)
				default:
					// Use the message from the HTTP error if available
					if he.Message != nil {
						if msg, ok := he.Message.(string); ok {
							err = NewWithMessage(ErrInternal, msg, err)
						} else {
							err = New(ErrInternal, err)
						}
					} else {
						err = New(ErrInternal, err)
					}
				}
			}
		}

		// Use our standard error response
		if err := RespondWithError(c, err); err != nil {
			// If responding with error fails, log it and send a basic response
			c.Logger().Error(err)
			_ = c.NoContent(http.StatusInternalServerError)
		}
	}
}