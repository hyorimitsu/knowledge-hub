package comment

import "github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"

// CreateCommentUseCase defines the interface for creating a comment
type CreateCommentUseCase interface {
	Execute(input CreateCommentInput) (*model.Comment, error)
}

// CreateCommentInput contains the data needed to create a comment
type CreateCommentInput struct {
	Content     string
	AuthorID    string
	KnowledgeID string
	TenantID    string
}

// UpdateCommentUseCase defines the interface for updating a comment
type UpdateCommentUseCase interface {
	Execute(input UpdateCommentInput) (*model.Comment, error)
}

// UpdateCommentInput contains the data needed to update a comment
type UpdateCommentInput struct {
	ID       string
	Content  string
	TenantID string
}

// DeleteCommentUseCase defines the interface for deleting a comment
type DeleteCommentUseCase interface {
	Execute(input DeleteCommentInput) error
}

// DeleteCommentInput contains the data needed to delete a comment
type DeleteCommentInput struct {
	ID       string
	TenantID string
}