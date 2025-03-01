package comment

import (
	"errors"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type deleteCommentUseCase struct {
	commentRepository repository.CommentRepository
	tenantRepository  repository.TenantRepository
}

// NewDeleteCommentUseCase creates a new instance of DeleteCommentUseCase
func NewDeleteCommentUseCase(
	commentRepository repository.CommentRepository,
	tenantRepository repository.TenantRepository,
) DeleteCommentUseCase {
	return &deleteCommentUseCase{
		commentRepository: commentRepository,
		tenantRepository:  tenantRepository,
	}
}

// Execute deletes a comment
func (uc *deleteCommentUseCase) Execute(input DeleteCommentInput) error {
	// Validate input
	if input.ID == "" {
		return errors.New("comment ID is required")
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

	// Find comment
	comment, err := uc.commentRepository.FindByID(input.ID, input.TenantID)
	if err != nil {
		return err
	}
	if comment == nil {
		return errors.New("comment not found")
	}

	// Delete comment
	return uc.commentRepository.Delete(input.ID, input.TenantID)
}