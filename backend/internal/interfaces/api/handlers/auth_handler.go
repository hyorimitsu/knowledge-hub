package handlers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	appErrors "github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/errors"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/user"
)

type AuthHandler struct {
	authenticateUserUseCase user.AuthenticateUserUseCase
	registerUserUseCase     user.RegisterUserUseCase
}

func NewAuthHandler(
	authenticateUserUseCase user.AuthenticateUserUseCase,
	registerUserUseCase user.RegisterUserUseCase,
) *AuthHandler {
	return &AuthHandler{
		authenticateUserUseCase: authenticateUserUseCase,
		registerUserUseCase:     registerUserUseCase,
	}
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	TenantID string `json:"tenant_id" validate:"required"`
}

// RegisterRequest represents the register request body
type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	TenantID string `json:"tenant_id" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=admin editor viewer"`
}

// TokenResponse represents the token response
type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

// Login handles user login
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return appErrors.NewValidationError("Invalid request body", nil, err)
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// Authenticate user
	user, err := h.authenticateUserUseCase.Execute(user.AuthenticateUserInput{
		Email:    req.Email,
		Password: req.Password,
		TenantID: req.TenantID,
	})
	if err != nil {
		return appErrors.Unauthorized("Invalid email or password", err)
	}

	// Generate JWT token
	token, expiresAt, err := generateToken(user.ID, user.Email, user.Role, user.TenantID)
	if err != nil {
		return appErrors.InternalServerError("Failed to generate token", err)
	}

	return appErrors.SendOK(c, TokenResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		UserID:    user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
	})
}

// Register handles user registration
// @Summary Register user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "User registration data"
// @Success 201 {object} TokenResponse
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 409 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return appErrors.NewValidationError("Invalid request body", nil, err)
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// Register user
	user, err := h.registerUserUseCase.Execute(user.RegisterUserInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		TenantID: req.TenantID,
		Role:     req.Role,
	})
	if err != nil {
		return appErrors.Conflict("User registration failed", err)
	}

	// Generate JWT token
	token, expiresAt, err := generateToken(user.ID, user.Email, user.Role, user.TenantID)
	if err != nil {
		return appErrors.InternalServerError("Failed to generate token", err)
	}

	return appErrors.SendCreated(c, TokenResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		UserID:    user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
	})
}

// generateToken generates a JWT token
func generateToken(userID, email, role, tenantID string) (string, int64, error) {
	// Get JWT secret from environment variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-key" // Default secret for development
	}

	// Set expiration time
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	// Create claims
	claims := &Claims{
		UserID:   userID,
		Email:    email,
		Role:     role,
		TenantID: tenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expiresAt, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "knowledge-hub",
			Subject:   userID,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresAt, nil
}

// Me handles getting the current user
// @Summary Get current user
// @Description Get the current authenticated user
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /auth/me [get]
func (h *AuthHandler) Me(c echo.Context) error {
	// Get user from context (set by auth middleware)
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims)

	return appErrors.SendOK(c, map[string]interface{}{
		"user_id":   claims.UserID,
		"email":     claims.Email,
		"role":      claims.Role,
		"tenant_id": claims.TenantID,
	})
}

// RegisterRoutes registers the public auth routes
func (h *AuthHandler) RegisterRoutes(g *echo.Group) {
	auth := g.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
}

// RegisterProtectedRoutes registers the protected auth routes
func (h *AuthHandler) RegisterProtectedRoutes(g *echo.Group) {
	auth := g.Group("/auth")
	auth.GET("/me", h.Me)
}