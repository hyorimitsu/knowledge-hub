package tag

import "github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"

// CreateTagUseCase defines the interface for creating a tag
type CreateTagUseCase interface {
	Execute(input CreateTagInput) (*model.Tag, error)
}

// CreateTagInput contains the data needed to create a tag
type CreateTagInput struct {
	Name     string
	Color    string
	TenantID string
}

// UpdateTagUseCase defines the interface for updating a tag
type UpdateTagUseCase interface {
	Execute(input UpdateTagInput) (*model.Tag, error)
}

// UpdateTagInput contains the data needed to update a tag
type UpdateTagInput struct {
	ID       string
	Name     string
	Color    string
	TenantID string
}

// DeleteTagUseCase defines the interface for deleting a tag
type DeleteTagUseCase interface {
	Execute(input DeleteTagInput) error
}

// DeleteTagInput contains the data needed to delete a tag
type DeleteTagInput struct {
	ID       string
	TenantID string
}