package errors

// ErrorCode represents a unique error code for each type of error
type ErrorCode string

// Error codes
const (
	// General errors
	ErrInternal        ErrorCode = "INTERNAL_ERROR"
	ErrNotFound        ErrorCode = "NOT_FOUND"
	ErrBadRequest      ErrorCode = "BAD_REQUEST"
	ErrUnauthorized    ErrorCode = "UNAUTHORIZED"
	ErrForbidden       ErrorCode = "FORBIDDEN"
	ErrConflict        ErrorCode = "CONFLICT"
	ErrValidation      ErrorCode = "VALIDATION_ERROR"
	ErrTimeout         ErrorCode = "TIMEOUT"
	ErrUnavailable     ErrorCode = "SERVICE_UNAVAILABLE"
	ErrNotImplemented  ErrorCode = "NOT_IMPLEMENTED"
	ErrTooManyRequests ErrorCode = "TOO_MANY_REQUESTS"

	// Domain specific errors
	ErrTenantNotFound     ErrorCode = "TENANT_NOT_FOUND"
	ErrUserNotFound       ErrorCode = "USER_NOT_FOUND"
	ErrKnowledgeNotFound  ErrorCode = "KNOWLEDGE_NOT_FOUND"
	ErrTagNotFound        ErrorCode = "TAG_NOT_FOUND"
	ErrCommentNotFound    ErrorCode = "COMMENT_NOT_FOUND"
	ErrInvalidCredentials ErrorCode = "INVALID_CREDENTIALS"
	ErrEmailAlreadyExists ErrorCode = "EMAIL_ALREADY_EXISTS"
	ErrDomainAlreadyExists ErrorCode = "DOMAIN_ALREADY_EXISTS"
	ErrInvalidRole        ErrorCode = "INVALID_ROLE"
)

// HTTPStatusCode maps error codes to HTTP status codes
func (e ErrorCode) HTTPStatusCode() int {
	switch e {
	case ErrNotFound, ErrTenantNotFound, ErrUserNotFound, ErrKnowledgeNotFound, ErrTagNotFound, ErrCommentNotFound:
		return 404 // Not Found
	case ErrBadRequest, ErrValidation:
		return 400 // Bad Request
	case ErrUnauthorized, ErrInvalidCredentials:
		return 401 // Unauthorized
	case ErrForbidden:
		return 403 // Forbidden
	case ErrConflict, ErrEmailAlreadyExists, ErrDomainAlreadyExists:
		return 409 // Conflict
	case ErrTimeout:
		return 408 // Request Timeout
	case ErrUnavailable:
		return 503 // Service Unavailable
	case ErrNotImplemented:
		return 501 // Not Implemented
	case ErrTooManyRequests:
		return 429 // Too Many Requests
	default:
		return 500 // Internal Server Error
	}
}

// DefaultMessage returns a default message for each error code
func (e ErrorCode) DefaultMessage() string {
	switch e {
	case ErrInternal:
		return "An internal error occurred"
	case ErrNotFound:
		return "Resource not found"
	case ErrBadRequest:
		return "Invalid request"
	case ErrUnauthorized:
		return "Authentication required"
	case ErrForbidden:
		return "Access denied"
	case ErrConflict:
		return "Resource conflict"
	case ErrValidation:
		return "Validation error"
	case ErrTimeout:
		return "Request timeout"
	case ErrUnavailable:
		return "Service unavailable"
	case ErrNotImplemented:
		return "Not implemented"
	case ErrTooManyRequests:
		return "Too many requests"
	case ErrTenantNotFound:
		return "Tenant not found"
	case ErrUserNotFound:
		return "User not found"
	case ErrKnowledgeNotFound:
		return "Knowledge not found"
	case ErrTagNotFound:
		return "Tag not found"
	case ErrCommentNotFound:
		return "Comment not found"
	case ErrInvalidCredentials:
		return "Invalid email or password"
	case ErrEmailAlreadyExists:
		return "Email already exists"
	case ErrDomainAlreadyExists:
		return "Domain already exists"
	case ErrInvalidRole:
		return "Invalid role"
	default:
		return "An error occurred"
	}
}