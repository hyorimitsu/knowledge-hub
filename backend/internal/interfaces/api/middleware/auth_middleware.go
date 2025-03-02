package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	appErrors "github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/errors"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/interfaces/api/handlers"
)

// JWTConfig returns the JWT middleware configuration
func JWTConfig() echojwt.Config {
	// Get JWT secret from environment variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-key" // Default secret for development
	}

	return echojwt.Config{
		SigningKey: []byte(jwtSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &handlers.Claims{}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return appErrors.Unauthorized("Invalid or expired token", err)
		},
	}
}

// RoleMiddleware returns a middleware that checks if the user has the required role
func RoleMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*handlers.Claims)

			// Check if user has one of the required roles
			hasRole := false
			for _, role := range roles {
				if claims.Role == role {
					hasRole = true
					break
				}
			}

			if !hasRole {
				return appErrors.Forbidden("Insufficient permissions", nil)
			}

			return next(c)
		}
	}
}

// TenantMiddleware returns a middleware that checks if the user belongs to the tenant
func TenantMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*handlers.Claims)

			// Get tenant ID from path parameter
			tenantID := c.Param("tenant_id")
			if tenantID != "" && tenantID != claims.TenantID {
				return appErrors.Forbidden("You don't have access to this tenant", nil)
			}

			return next(c)
		}
	}
}