package user

import "github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"

// RegisterUserUseCase defines the interface for registering a user
type RegisterUserUseCase interface {
	Execute(input RegisterUserInput) (*model.User, error)
}

// RegisterUserInput contains the data needed to register a user
type RegisterUserInput struct {
	Name     string
	Email    string
	Password string
	TenantID string
	Role     string
}

// AuthenticateUserUseCase defines the interface for authenticating a user
type AuthenticateUserUseCase interface {
	Execute(input AuthenticateUserInput) (*model.User, error)
}

// AuthenticateUserInput contains the data needed to authenticate a user
type AuthenticateUserInput struct {
	Email    string
	Password string
	TenantID string
}

// UpdateUserUseCase defines the interface for updating user information
type UpdateUserUseCase interface {
	Execute(input UpdateUserInput) (*model.User, error)
}

// UpdateUserInput contains the data needed to update a user
type UpdateUserInput struct {
	ID       string
	Name     string
	Email    string
	Password string
	Avatar   string
	TenantID string
	Role     string
}