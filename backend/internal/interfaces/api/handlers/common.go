package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

// Claims represents the JWT claims
type Claims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	TenantID string `json:"tenant_id"`
	jwt.RegisteredClaims
}

// RepositoriesProvider defines the interface for accessing repositories
type RepositoriesProvider interface {
	Tenant() repository.TenantRepository
	User() repository.UserRepository
	Knowledge() repository.KnowledgeRepository
	Tag() repository.TagRepository
	Comment() repository.CommentRepository
}

// getUserClaims extracts the user claims from the context
func getUserClaims(c echo.Context) *Claims {
	user := c.Get("user")
	if user == nil {
		return nil
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return nil
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil
	}

	return claims
}