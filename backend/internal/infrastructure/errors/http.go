package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTP error helper functions

// BadRequest returns a 400 Bad Request error
func BadRequest(message string, err error) error {
	return NewWithMessage(ErrBadRequest, message, err)
}

// Unauthorized returns a 401 Unauthorized error
func Unauthorized(message string, err error) error {
	if message == "" {
		message = "Authentication required"
	}
	return NewWithMessage(ErrUnauthorized, message, err)
}

// Forbidden returns a 403 Forbidden error
func Forbidden(message string, err error) error {
	if message == "" {
		message = "Access denied"
	}
	return NewWithMessage(ErrForbidden, message, err)
}

// NotFound returns a 404 Not Found error
func NotFound(message string, err error) error {
	if message == "" {
		message = "Resource not found"
	}
	return NewWithMessage(ErrNotFound, message, err)
}

// Conflict returns a 409 Conflict error
func Conflict(message string, err error) error {
	if message == "" {
		message = "Resource conflict"
	}
	return NewWithMessage(ErrConflict, message, err)
}

// InternalServerError returns a 500 Internal Server Error
func InternalServerError(message string, err error) error {
	if message == "" {
		message = "Internal server error"
	}
	return NewWithMessage(ErrInternal, message, err)
}

// NotImplemented returns a 501 Not Implemented error
func NotImplemented(message string, err error) error {
	if message == "" {
		message = "Not implemented"
	}
	return NewWithMessage(ErrNotImplemented, message, err)
}

// ServiceUnavailable returns a 503 Service Unavailable error
func ServiceUnavailable(message string, err error) error {
	if message == "" {
		message = "Service unavailable"
	}
	return NewWithMessage(ErrUnavailable, message, err)
}

// TooManyRequests returns a 429 Too Many Requests error
func TooManyRequests(message string, err error) error {
	if message == "" {
		message = "Too many requests"
	}
	return NewWithMessage(ErrTooManyRequests, message, err)
}

// Domain-specific error helpers

// TenantNotFound returns a tenant not found error
func TenantNotFound(err error) error {
	return NewWithMessage(ErrTenantNotFound, "Tenant not found", err)
}

// UserNotFound returns a user not found error
func UserNotFound(err error) error {
	return NewWithMessage(ErrUserNotFound, "User not found", err)
}

// KnowledgeNotFound returns a knowledge not found error
func KnowledgeNotFound(err error) error {
	return NewWithMessage(ErrKnowledgeNotFound, "Knowledge not found", err)
}

// TagNotFound returns a tag not found error
func TagNotFound(err error) error {
	return NewWithMessage(ErrTagNotFound, "Tag not found", err)
}

// CommentNotFound returns a comment not found error
func CommentNotFound(err error) error {
	return NewWithMessage(ErrCommentNotFound, "Comment not found", err)
}

// InvalidCredentials returns an invalid credentials error
func InvalidCredentials(err error) error {
	return NewWithMessage(ErrInvalidCredentials, "Invalid email or password", err)
}

// EmailAlreadyExists returns an email already exists error
func EmailAlreadyExists(err error) error {
	return NewWithMessage(ErrEmailAlreadyExists, "Email already exists", err)
}

// DomainAlreadyExists returns a domain already exists error
func DomainAlreadyExists(err error) error {
	return NewWithMessage(ErrDomainAlreadyExists, "Domain already exists", err)
}

// InvalidRole returns an invalid role error
func InvalidRole(err error) error {
	return NewWithMessage(ErrInvalidRole, "Invalid role", err)
}

// HTTP response helpers

// SendOK sends a 200 OK response
func SendOK(c echo.Context, data interface{}) error {
	return RespondWithSuccess(c, http.StatusOK, data)
}

// SendCreated sends a 201 Created response
func SendCreated(c echo.Context, data interface{}) error {
	return RespondWithSuccess(c, http.StatusCreated, data)
}

// SendNoContent sends a 204 No Content response
func SendNoContent(c echo.Context) error {
	return RespondWithEmpty(c, http.StatusNoContent)
}

// SendAccepted sends a 202 Accepted response
func SendAccepted(c echo.Context, data interface{}) error {
	return RespondWithSuccess(c, http.StatusAccepted, data)
}