package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type updateUserUseCase struct {
	userRepository   repository.UserRepository
	tenantRepository repository.TenantRepository
}

// NewUpdateUserUseCase creates a new instance of UpdateUserUseCase
func NewUpdateUserUseCase(
	userRepository repository.UserRepository,
	tenantRepository repository.TenantRepository,
) UpdateUserUseCase {
	return &updateUserUseCase{
		userRepository:   userRepository,
		tenantRepository: tenantRepository,
	}
}

// Execute updates a user
func (uc *updateUserUseCase) Execute(input UpdateUserInput) (*model.User, error) {
	// Validate input
	if input.ID == "" {
		return nil, errors.New("user ID is required")
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

	// Find user
	user, err := uc.userRepository.FindByID(input.ID, input.TenantID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Check if email is being changed and if it's already in use
	if input.Email != "" && input.Email != user.Email {
		existingUser, err := uc.userRepository.FindByEmail(input.Email, input.TenantID)
		if err == nil && existingUser != nil && existingUser.ID != input.ID {
			return nil, errors.New("email already in use")
		}
		user.Email = input.Email
	}

	// Update user fields if provided
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if input.Avatar != "" {
		user.Avatar = input.Avatar
	}
	if input.Role != "" {
		role := model.Role(input.Role)
		if !role.IsValid() {
			return nil, errors.New("invalid role")
		}
		user.Role = role.String()
	}

	user.UpdatedAt = time.Now()

	// Save user
	err = uc.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	// Return user without password
	user.Password = ""
	return user, nil
}