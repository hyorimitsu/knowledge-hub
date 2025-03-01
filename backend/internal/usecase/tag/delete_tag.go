package tag

import (
	"errors"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type deleteTagUseCase struct {
	tagRepository    repository.TagRepository
	tenantRepository repository.TenantRepository
}

// NewDeleteTagUseCase creates a new instance of DeleteTagUseCase
func NewDeleteTagUseCase(
	tagRepository repository.TagRepository,
	tenantRepository repository.TenantRepository,
) DeleteTagUseCase {
	return &deleteTagUseCase{
		tagRepository:    tagRepository,
		tenantRepository: tenantRepository,
	}
}

// Execute deletes a tag
func (uc *deleteTagUseCase) Execute(input DeleteTagInput) error {
	// Validate input
	if input.ID == "" {
		return errors.New("tag ID is required")
	}
	if input.TenantID == "" {
		return errors.New("tenant ID is required")
	}

	// Verify tenant exists
	tenant, err := uc.tenantRepository.FindByID(input.TenantID)
	if err != nil {
		return err
	}
	if tenant == nil {
		return errors.New("tenant not found")
	}

	// Find tag
	tag, err := uc.tagRepository.FindByID(input.ID, input.TenantID)
	if err != nil {
		return err
	}
	if tag == nil {
		return errors.New("tag not found")
	}

	// Delete tag
	return uc.tagRepository.Delete(input.ID, input.TenantID)
}