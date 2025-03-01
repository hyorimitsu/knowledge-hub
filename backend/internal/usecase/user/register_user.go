package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type registerUserUseCase struct {
	userRepository   repository.UserRepository
	tenantRepository repository.TenantRepository
}

// NewRegisterUserUseCase creates a new instance of RegisterUserUseCase
func NewRegisterUserUseCase(
	userRepository repository.UserRepository,
	tenantRepository repository.TenantRepository,
) RegisterUserUseCase {
	return &registerUserUseCase{
		userRepository:   userRepository,
		tenantRepository: tenantRepository,
	}
}

// Execute registers a new user
func (uc *registerUserUseCase) Execute(input RegisterUserInput) (*model.User, error) {
	// Validate input
	if input.Name == "" {
		return nil, errors.New("user name is required")
	}
	if input.Email == "" {
		return nil, errors.New("user email is required")
	}
	if input.Password == "" {
		return nil, errors.New("user password is required")
	}
	if input.TenantID == "" {
		return nil, errors.New("tenant ID is required")
	}

	// Validate role
	role := model.Role(input.Role)
	if !role.IsValid() {
		return nil, errors.New("invalid role")
	}

	// Verify tenant exists
	tenant, err := uc.tenantRepository.FindByID(input.TenantID)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, errors.New("tenant not found")
	}

	// Check if email already exists for this tenant
	existingUser, err := uc.userRepository.FindByEmail(input.Email, input.TenantID)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered for this tenant")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	now := time.Now()
	user := &model.User{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		TenantID:  input.TenantID,
		Role:      role.String(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Save user
	err = uc.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}