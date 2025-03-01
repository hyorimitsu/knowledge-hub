package knowledge

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type searchKnowledgeUseCase struct {
	knowledgeRepository repository.KnowledgeRepository
	tagRepository       repository.TagRepository
	tenantRepository    repository.TenantRepository
	db                  *gorm.DB
}

// NewSearchKnowledgeUseCase creates a new instance of SearchKnowledgeUseCase
func NewSearchKnowledgeUseCase(
	knowledgeRepository repository.KnowledgeRepository,
	tagRepository repository.TagRepository,
	tenantRepository repository.TenantRepository,
	db *gorm.DB,
) SearchKnowledgeUseCase {
	return &searchKnowledgeUseCase{
		knowledgeRepository: knowledgeRepository,
		tagRepository:       tagRepository,
		tenantRepository:    tenantRepository,
		db:                  db,
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

	// Build query
	query := uc.db.Model(&model.Knowledge{}).
		Preload("Tags").
		Preload("Comments").
		Where("tenant_id = ?", input.TenantID)

	// Add search conditions
	if input.Query != "" {
		searchQuery := "%" + strings.ToLower(input.Query) + "%"
		query = query.Where("LOWER(title) LIKE ? OR LOWER(content) LIKE ?", searchQuery, searchQuery)
	}

	// Filter by author if provided
	if input.AuthorID != "" {
		query = query.Where("author_id = ?", input.AuthorID)
	}

	// Filter by tags if provided
	if len(input.TagIDs) > 0 {
		query = query.Joins("JOIN knowledge_tags ON knowledge_tags.knowledge_id = knowledges.id").
			Where("knowledge_tags.tag_id IN ?", input.TagIDs).
			Group("knowledges.id")
	}

	// Execute query
	var results []*model.Knowledge
	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}