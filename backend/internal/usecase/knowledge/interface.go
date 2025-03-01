package knowledge

import "github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"

// CreateKnowledgeUseCase defines the interface for creating knowledge
type CreateKnowledgeUseCase interface {
	Execute(input CreateKnowledgeInput) (*model.Knowledge, error)
}

// CreateKnowledgeInput contains the data needed to create knowledge
type CreateKnowledgeInput struct {
	Title    string
	Content  string
	AuthorID string
	TenantID string
	TagIDs   []string
}

// UpdateKnowledgeUseCase defines the interface for updating knowledge
type UpdateKnowledgeUseCase interface {
	Execute(input UpdateKnowledgeInput) (*model.Knowledge, error)
}

// UpdateKnowledgeInput contains the data needed to update knowledge
type UpdateKnowledgeInput struct {
	ID       string
	Title    string
	Content  string
	TenantID string
	TagIDs   []string
}

// DeleteKnowledgeUseCase defines the interface for deleting knowledge
type DeleteKnowledgeUseCase interface {
	Execute(input DeleteKnowledgeInput) error
}

// DeleteKnowledgeInput contains the data needed to delete knowledge
type DeleteKnowledgeInput struct {
	ID       string
	TenantID string
}

// SearchKnowledgeUseCase defines the interface for searching knowledge
type SearchKnowledgeUseCase interface {
	Execute(input SearchKnowledgeInput) ([]*model.Knowledge, error)
}

// SearchKnowledgeInput contains the data needed to search knowledge
type SearchKnowledgeInput struct {
	Query    string
	TenantID string
	TagIDs   []string
	AuthorID string
}