package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type authenticateUserUseCase struct {
	userRepository   repository.UserRepository
	tenantRepository repository.TenantRepository
}

// NewAuthenticateUserUseCase creates a new instance of AuthenticateUserUseCase
func NewAuthenticateUserUseCase(
	userRepository repository.UserRepository,
	tenantRepository repository.TenantRepository,
) AuthenticateUserUseCase {
	return &authenticateUserUseCase{
		userRepository:   userRepository,
		tenantRepository: tenantRepository,
	}
}

// Execute authenticates a user
func (uc *authenticateUserUseCase) Execute(input AuthenticateUserInput) (*model.User, error) {
	// Validate input
	if input.Email == "" {
		return nil, errors.New("email is required")
	}
	if input.Password == "" {
		return nil, errors.New("password is required")
	}
	if input.TenantID == "" {
		return nil, errors.New("tenant ID is required")
	}

	// Verify tenant exists
	tenant, err := uc.tenantRepository.FindByID(input.TenantID)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, errors.New("tenant not found")
	}

	// Find user by email
	user, err := uc.userRepository.FindByEmail(input.Email, input.TenantID)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Return user without password
	user.Password = ""
	return user, nil
}