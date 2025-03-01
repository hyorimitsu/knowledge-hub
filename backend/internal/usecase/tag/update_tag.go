package tag

import (
	"errors"
	"time"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type updateTagUseCase struct {
	tagRepository    repository.TagRepository
	tenantRepository repository.TenantRepository
}

// NewUpdateTagUseCase creates a new instance of UpdateTagUseCase
func NewUpdateTagUseCase(
	tagRepository repository.TagRepository,
	tenantRepository repository.TenantRepository,
) UpdateTagUseCase {
	return &updateTagUseCase{
		tagRepository:    tagRepository,
		tenantRepository: tenantRepository,
	}
}

// Execute updates a tag
func (uc *updateTagUseCase) Execute(input UpdateTagInput) (*model.Tag, error) {
	// Validate input
	if input.ID == "" {
		return nil, errors.New("tag ID is required")
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

	// Find tag
	tag, err := uc.tagRepository.FindByID(input.ID, input.TenantID)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, errors.New("tag not found")
	}

	// Update tag fields if provided
	if input.Name != "" {
		tag.Name = input.Name
	}
	if input.Color != "" {
		tag.Color = input.Color
	}

	tag.UpdatedAt = time.Now()

	// Save tag
	err = uc.tagRepository.Update(tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}