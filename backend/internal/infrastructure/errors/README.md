# Error Handling System

This package provides a comprehensive error handling system for the Knowledge Hub backend application. It includes custom error types, error codes, standardized error responses, validation error handling, and logging functionality.

## Features

- Custom error types for different layers (domain, application, infrastructure)
- Error codes with HTTP status code mapping
- Standardized error response format
- Validation error handling
- Error logging with stack traces
- Echo middleware for error handling
- Helper functions for common HTTP errors

## Error Response Format

All API responses follow a consistent format:

### Success Response

```json
{
  "success": true,
  "data": {
    // Response data
  }
}
```

### Error Response

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": {
      // Optional error details
    }
  }
}
```

## Usage Examples

### Creating and Returning Errors

```go
// Basic error
return errors.New(errors.ErrNotFound, err)

// Error with custom message
return errors.NewWithMessage(errors.ErrNotFound, "User not found", err)

// Error with details
details := map[string]interface{}{
    "userId": "123",
    "reason": "Account deleted",
}
return errors.NewWithDetails(errors.ErrNotFound, "User not found", details, err)

// Using helper functions
return errors.NotFound("User not found", err)
return errors.BadRequest("Invalid input", err)
return errors.Unauthorized("Authentication required", err)
```

### Validation Errors

```go
// Creating validation errors manually
fieldErrors := map[string]string{
    "username": "Username is required",
    "email":    "Email must be valid",
}
return errors.NewValidationError("Validation failed", fieldErrors, err)

// Using the validator
type CreateUserRequest struct {
    Username string `json:"username" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"min=18"`
}

func CreateUserHandler(c echo.Context) error {
    var req CreateUserRequest
    if err := errors.BindAndValidate(c, &req); err != nil {
        return err // Automatically returns validation error response
    }
    
    // Process the request...
    return errors.SendCreated(c, user)
}
```

### Success Responses

```go
// 200 OK with data
return errors.SendOK(c, data)

// 201 Created with data
return errors.SendCreated(c, data)

// 204 No Content
return errors.SendNoContent(c)
```

### Domain-Specific Errors

```go
// Domain error
return errors.NewDomainError(errors.ErrTenantNotFound, "Tenant not found", err)

// Application error
return errors.NewApplicationError(errors.ErrInvalidCredentials, "Invalid email or password", err)

// Infrastructure error
return errors.NewInfrastructureError(errors.ErrUnavailable, "Database connection failed", err)
```

### Logging

```go
logger := errors.NewLogger(errors.DefaultLogConfig)

// Log messages
logger.Debug("Debug message")
logger.Info("Info message")
logger.Warn("Warning message")
logger.Error("Error message")

// Log error with stack trace
logger.ErrorWithStack(err)
```

## Error Codes

The system defines the following error codes:

### General Errors
- `INTERNAL_ERROR`: Internal server error (500)
- `NOT_FOUND`: Resource not found (404)
- `BAD_REQUEST`: Invalid request (400)
- `UNAUTHORIZED`: Authentication required (401)
- `FORBIDDEN`: Access denied (403)
- `CONFLICT`: Resource conflict (409)
- `VALIDATION_ERROR`: Validation error (400)
- `TIMEOUT`: Request timeout (408)
- `SERVICE_UNAVAILABLE`: Service unavailable (503)
- `NOT_IMPLEMENTED`: Not implemented (501)
- `TOO_MANY_REQUESTS`: Too many requests (429)

### Domain-Specific Errors
- `TENANT_NOT_FOUND`: Tenant not found (404)
- `USER_NOT_FOUND`: User not found (404)
- `KNOWLEDGE_NOT_FOUND`: Knowledge not found (404)
- `TAG_NOT_FOUND`: Tag not found (404)
- `COMMENT_NOT_FOUND`: Comment not found (404)
- `INVALID_CREDENTIALS`: Invalid credentials (401)
- `EMAIL_ALREADY_EXISTS`: Email already exists (409)
- `DOMAIN_ALREADY_EXISTS`: Domain already exists (409)
- `INVALID_ROLE`: Invalid role (400)