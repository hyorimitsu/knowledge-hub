package knowledge

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type createKnowledgeUseCase struct {
	knowledgeRepository repository.KnowledgeRepository
	userRepository      repository.UserRepository
	tagRepository       repository.TagRepository
	tenantRepository    repository.TenantRepository
}

// NewCreateKnowledgeUseCase creates a new instance of CreateKnowledgeUseCase
func NewCreateKnowledgeUseCase(
	knowledgeRepository repository.KnowledgeRepository,
	userRepository repository.UserRepository,
	tagRepository repository.TagRepository,
	tenantRepository repository.TenantRepository,
) CreateKnowledgeUseCase {
	return &createKnowledgeUseCase{
		knowledgeRepository: knowledgeRepository,
		userRepository:      userRepository,
		tagRepository:       tagRepository,
		tenantRepository:    tenantRepository,
	}
}

// Execute creates a new knowledge
func (uc *createKnowledgeUseCase) Execute(input CreateKnowledgeInput) (*model.Knowledge, error) {
	// Validate input
	if input.Title == "" {
		return nil, errors.New("title is required")
	}
	if input.Content == "" {
		return nil, errors.New("content is required")
	}
	if input.AuthorID == "" {
		return nil, errors.New("author ID is required")
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

	// Verify author exists
	author, err := uc.userRepository.FindByID(input.AuthorID, input.TenantID)
	if err != nil {
		return nil, err
	}
	if author == nil {
		return nil, errors.New("author not found")
	}

	// Create knowledge
	now := time.Now()
	knowledge := &model.Knowledge{
		ID:        uuid.New().String(),
		Title:     input.Title,
		Content:   input.Content,
		AuthorID:  input.AuthorID,
		TenantID:  input.TenantID,
		Tags:      []model.Tag{},
		Comments:  []model.Comment{},
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Add tags if provided
	if len(input.TagIDs) > 0 {
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

	// Save knowledge
	err = uc.knowledgeRepository.Create(knowledge)
	if err != nil {
		return nil, err
	}

	return knowledge, nil
}