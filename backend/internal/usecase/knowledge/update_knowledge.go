package knowledge

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type updateKnowledgeUseCase struct {
	knowledgeRepository repository.KnowledgeRepository
	tagRepository       repository.TagRepository
	tenantRepository    repository.TenantRepository
}

// NewUpdateKnowledgeUseCase creates a new instance of UpdateKnowledgeUseCase
func NewUpdateKnowledgeUseCase(
	knowledgeRepository repository.KnowledgeRepository,
	tagRepository repository.TagRepository,
	tenantRepository repository.TenantRepository,
) UpdateKnowledgeUseCase {
	return &updateKnowledgeUseCase{
		knowledgeRepository: knowledgeRepository,
		tagRepository:       tagRepository,
		tenantRepository:    tenantRepository,
	}
}

// Execute updates knowledge
func (uc *updateKnowledgeUseCase) Execute(input UpdateKnowledgeInput) (*model.Knowledge, error) {
	// Validate input
	if input.ID == "" {
		return nil, errors.New("knowledge ID is required")
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

	// Find knowledge
	knowledge, err := uc.knowledgeRepository.FindByID(input.ID, input.TenantID)
	if err != nil {
		return nil, err
	}
	if knowledge == nil {
		return nil, errors.New("knowledge not found")
	}

	// Update knowledge fields if provided
	if input.Title != "" {
		knowledge.Title = input.Title
	}
	if input.Content != "" {
		knowledge.Content = input.Content
	}

	// Update tags if provided
	if len(input.TagIDs) > 0 {
		knowledge.Tags = []model.Tag{}
		for _, tagID := range input.TagIDs {
			tag, err := uc.tagRepository.FindByID(tagID, input.TenantID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					continue
				}
				return nil, err
			}
			if tag != nil {
				knowledge.Tags = append(knowledge.Tags, *tag)
			}
		}
	}

	knowledge.UpdatedAt = time.Now()

	// Save knowledge
	err = uc.knowledgeRepository.Update(knowledge)
	if err != nil {
		return nil, err
	}

	return knowledge, nil
}