package knowledge

import (
	"errors"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type deleteKnowledgeUseCase struct {
	knowledgeRepository repository.KnowledgeRepository
	tenantRepository    repository.TenantRepository
}

// NewDeleteKnowledgeUseCase creates a new instance of DeleteKnowledgeUseCase
func NewDeleteKnowledgeUseCase(
	knowledgeRepository repository.KnowledgeRepository,
	tenantRepository repository.TenantRepository,
) DeleteKnowledgeUseCase {
	return &deleteKnowledgeUseCase{
		knowledgeRepository: knowledgeRepository,
		tenantRepository:    tenantRepository,
	}
}

// Execute deletes knowledge
func (uc *deleteKnowledgeUseCase) Execute(input DeleteKnowledgeInput) error {
	// Validate input
	if input.ID == "" {
		return errors.New("knowledge ID is required")
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

	// Find knowledge
	knowledge, err := uc.knowledgeRepository.FindByID(input.ID, input.TenantID)
	if err != nil {
		return err
	}
	if knowledge == nil {
		return errors.New("knowledge not found")
	}

	// Delete knowledge
	return uc.knowledgeRepository.Delete(input.ID, input.TenantID)
}