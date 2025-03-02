package knowledge

import (
	"errors"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type searchKnowledgeUseCase struct {
	knowledgeRepository repository.KnowledgeRepository
	tagRepository       repository.TagRepository
	tenantRepository    repository.TenantRepository
}

// NewSearchKnowledgeUseCase creates a new instance of SearchKnowledgeUseCase
func NewSearchKnowledgeUseCase(
	knowledgeRepository repository.KnowledgeRepository,
	tagRepository repository.TagRepository,
	tenantRepository repository.TenantRepository,
) SearchKnowledgeUseCase {
	return &searchKnowledgeUseCase{
		knowledgeRepository: knowledgeRepository,
		tagRepository:       tagRepository,
		tenantRepository:    tenantRepository,
	}
}

// Execute searches for knowledge
func (uc *searchKnowledgeUseCase) Execute(input SearchKnowledgeInput) ([]*model.Knowledge, error) {
	// Validate input
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

	// Use the repository's search method
	results, err := uc.knowledgeRepository.Search(input.Query, input.TenantID, input.TagIDs, input.AuthorID)
	if err != nil {
		return nil, err
	}

	return results, nil
}