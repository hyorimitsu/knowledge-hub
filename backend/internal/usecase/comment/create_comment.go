package comment

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type createCommentUseCase struct {
	commentRepository   repository.CommentRepository
	knowledgeRepository repository.KnowledgeRepository
	userRepository      repository.UserRepository
	tenantRepository    repository.TenantRepository
}

// NewCreateCommentUseCase creates a new instance of CreateCommentUseCase
func NewCreateCommentUseCase(
	commentRepository repository.CommentRepository,
	knowledgeRepository repository.KnowledgeRepository,
	userRepository repository.UserRepository,
	tenantRepository repository.TenantRepository,
) CreateCommentUseCase {
	return &createCommentUseCase{
		commentRepository:   commentRepository,
		knowledgeRepository: knowledgeRepository,
		userRepository:      userRepository,
		tenantRepository:    tenantRepository,
	}
}

// Execute creates a new comment
func (uc *createCommentUseCase) Execute(input CreateCommentInput) (*model.Comment, error) {
	// Validate input
	if input.Content == "" {
		return nil, errors.New("comment content is required")
	}
	if input.AuthorID == "" {
		return nil, errors.New("author ID is required")
	}
	if input.KnowledgeID == "" {
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

	// Verify knowledge exists
	knowledge, err := uc.knowledgeRepository.FindByID(input.KnowledgeID, input.TenantID)
	if err != nil {
		return nil, err
	}
	if knowledge == nil {
		return nil, errors.New("knowledge not found")
	}

	// Verify author exists
	author, err := uc.userRepository.FindByID(input.AuthorID, input.TenantID)
	if err != nil {
		return nil, err
	}
	if author == nil {
		return nil, errors.New("author not found")
	}

	// Create comment
	now := time.Now()
	comment := &model.Comment{
		ID:          uuid.New().String(),
		Content:     input.Content,
		AuthorID:    input.AuthorID,
		KnowledgeID: input.KnowledgeID,
		TenantID:    input.TenantID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Save comment
	err = uc.commentRepository.Create(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}