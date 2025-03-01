package comment

import (
	"errors"
	"time"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type updateCommentUseCase struct {
	commentRepository repository.CommentRepository
	tenantRepository  repository.TenantRepository
}

// NewUpdateCommentUseCase creates a new instance of UpdateCommentUseCase
func NewUpdateCommentUseCase(
	commentRepository repository.CommentRepository,
	tenantRepository repository.TenantRepository,
) UpdateCommentUseCase {
	return &updateCommentUseCase{
		commentRepository: commentRepository,
		tenantRepository:  tenantRepository,
	}
}

// Execute updates a comment
func (uc *updateCommentUseCase) Execute(input UpdateCommentInput) (*model.Comment, error) {
	// Validate input
	if input.ID == "" {
		return nil, errors.New("comment ID is required")
	}
	if input.Content == "" {
		return nil, errors.New("comment content is required")
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

	// Find comment
	comment, err := uc.commentRepository.FindByID(input.ID, input.TenantID)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, errors.New("comment not found")
	}

	// Update comment
	comment.Content = input.Content
	comment.UpdatedAt = time.Now()

	// Save comment
	err = uc.commentRepository.Update(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}