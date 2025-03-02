package tag

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type createTagUseCase struct {
	tagRepository    repository.TagRepository
	tenantRepository repository.TenantRepository
}

// NewCreateTagUseCase creates a new instance of CreateTagUseCase
func NewCreateTagUseCase(
	tagRepository repository.TagRepository,
	tenantRepository repository.TenantRepository,
) CreateTagUseCase {
	return &createTagUseCase{
		tagRepository:    tagRepository,
		tenantRepository: tenantRepository,
	}
}

// Execute creates a new tag
func (uc *createTagUseCase) Execute(input CreateTagInput) (*model.Tag, error) {
	// Validate input
	if input.Name == "" {
		return nil, errors.New("tag name is required")
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

	// Create tag
	now := time.Now()
	tag := &model.Tag{
		ID:        uuid.New().String(),
		Name:      input.Name,
		TenantID:  input.TenantID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Save tag
	err = uc.tagRepository.Create(tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}