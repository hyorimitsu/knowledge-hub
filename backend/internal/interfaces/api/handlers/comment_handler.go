package handlers

import (
	"github.com/labstack/echo/v4"

	appErrors "github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/errors"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/comment"
)

type CommentHandler struct {
	createCommentUseCase comment.CreateCommentUseCase
	updateCommentUseCase comment.UpdateCommentUseCase
	deleteCommentUseCase comment.DeleteCommentUseCase
}

func NewCommentHandler(
	createCommentUseCase comment.CreateCommentUseCase,
	updateCommentUseCase comment.UpdateCommentUseCase,
	deleteCommentUseCase comment.DeleteCommentUseCase,
) *CommentHandler {
	return &CommentHandler{
		createCommentUseCase: createCommentUseCase,
		updateCommentUseCase: updateCommentUseCase,
		deleteCommentUseCase: deleteCommentUseCase,
	}
}

// CreateCommentRequest represents the create comment request body
type CreateCommentRequest struct {
	Content string `json:"content" validate:"required"`
}

// UpdateCommentRequest represents the update comment request body
type UpdateCommentRequest struct {
	Content string `json:"content" validate:"required"`
}

// Create handles creating a new comment
// @Summary Create comment
// @Description Create a new comment for a knowledge
// @Tags comments
// @Accept json
// @Produce json
// @Param knowledge_id path string true "Knowledge ID"
// @Param request body CreateCommentRequest true "Comment data"
// @Security ApiKeyAuth
// @Success 201 {object} model.Comment
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge/{knowledge_id}/comments [post]
func (h *CommentHandler) Create(c echo.Context) error {
	knowledgeID := c.Param("knowledge_id")
	if knowledgeID == "" {
		return appErrors.NewValidationError("Knowledge ID is required", nil, nil)
	}

	var req CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		return appErrors.NewValidationError("Invalid request body", nil, err)
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Create comment
	comment, err := h.createCommentUseCase.Execute(comment.CreateCommentInput{
		Content:     req.Content,
		AuthorID:    claims.UserID,
		KnowledgeID: knowledgeID,
		TenantID:    claims.TenantID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to create comment", err)
	}

	return appErrors.SendCreated(c, comment)
}

// Update handles updating a comment
// @Summary Update comment
// @Description Update a comment
// @Tags comments
// @Accept json
// @Produce json
// @Param knowledge_id path string true "Knowledge ID"
// @Param comment_id path string true "Comment ID"
// @Param request body UpdateCommentRequest true "Comment data"
// @Security ApiKeyAuth
// @Success 200 {object} model.Comment
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 403 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge/{knowledge_id}/comments/{comment_id} [put]
func (h *CommentHandler) Update(c echo.Context) error {
	knowledgeID := c.Param("knowledge_id")
	if knowledgeID == "" {
		return appErrors.NewValidationError("Knowledge ID is required", nil, nil)
	}

	commentID := c.Param("comment_id")
	if commentID == "" {
		return appErrors.NewValidationError("Comment ID is required", nil, nil)
	}

	var req UpdateCommentRequest
	if err := c.Bind(&req); err != nil {
		return appErrors.NewValidationError("Invalid request body", nil, err)
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Get repositories
	repo := c.Get("repositories").(RepositoriesProvider).Comment()

	// Check if comment exists and belongs to the user
	existingComment, err := repo.FindByID(commentID, claims.TenantID)
	if err != nil {
		return appErrors.NotFound("Comment not found", err)
	}

	// Check if user is the author of the comment
	if existingComment.AuthorID != claims.UserID && claims.Role != "admin" {
		return appErrors.Forbidden("You don't have permission to update this comment", nil)
	}

	// Update comment
	comment, err := h.updateCommentUseCase.Execute(comment.UpdateCommentInput{
		ID:       commentID,
		Content:  req.Content,
		TenantID: claims.TenantID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to update comment", err)
	}

	return appErrors.SendOK(c, comment)
}

// Delete handles deleting a comment
// @Summary Delete comment
// @Description Delete a comment
// @Tags comments
// @Accept json
// @Produce json
// @Param knowledge_id path string true "Knowledge ID"
// @Param comment_id path string true "Comment ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 403 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge/{knowledge_id}/comments/{comment_id} [delete]
func (h *CommentHandler) Delete(c echo.Context) error {
	knowledgeID := c.Param("knowledge_id")
	if knowledgeID == "" {
		return appErrors.NewValidationError("Knowledge ID is required", nil, nil)
	}

	commentID := c.Param("comment_id")
	if commentID == "" {
		return appErrors.NewValidationError("Comment ID is required", nil, nil)
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Get repositories
	repo := c.Get("repositories").(RepositoriesProvider).Comment()

	// Check if comment exists and belongs to the user
	existingComment, err := repo.FindByID(commentID, claims.TenantID)
	if err != nil {
		return appErrors.NotFound("Comment not found", err)
	}

	// Check if user is the author of the comment or an admin
	if existingComment.AuthorID != claims.UserID && claims.Role != "admin" {
		return appErrors.Forbidden("You don't have permission to delete this comment", nil)
	}

	// Delete comment
	err = h.deleteCommentUseCase.Execute(comment.DeleteCommentInput{
		ID:       commentID,
		TenantID: claims.TenantID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to delete comment", err)
	}

	return appErrors.SendNoContent(c)
}

// List handles listing comments for a knowledge
// @Summary List comments
// @Description List all comments for a knowledge
// @Tags comments
// @Accept json
// @Produce json
// @Param knowledge_id path string true "Knowledge ID"
// @Security ApiKeyAuth
// @Success 200 {array} model.Comment
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge/{knowledge_id}/comments [get]
func (h *CommentHandler) List(c echo.Context) error {
	knowledgeID := c.Param("knowledge_id")
	if knowledgeID == "" {
		return appErrors.NewValidationError("Knowledge ID is required", nil, nil)
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Get repositories
	repo := c.Get("repositories").(RepositoriesProvider).Comment()

	// List comments
	comments, err := repo.FindByKnowledgeID(knowledgeID, claims.TenantID)
	if err != nil {
		return appErrors.InternalServerError("Failed to list comments", err)
	}

	return appErrors.SendOK(c, comments)
}

// RegisterRoutes registers the comment routes
func (h *CommentHandler) RegisterRoutes(g *echo.Group) {
	knowledge := g.Group("/knowledge/:knowledge_id/comments")
	knowledge.POST("", h.Create)
	knowledge.GET("", h.List)
	knowledge.PUT("/:comment_id", h.Update)
	knowledge.DELETE("/:comment_id", h.Delete)
}