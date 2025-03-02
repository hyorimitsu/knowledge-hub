package handlers

import (
	"github.com/labstack/echo/v4"

	appErrors "github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/errors"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/knowledge"
)

type KnowledgeHandler struct {
	createKnowledgeUseCase knowledge.CreateKnowledgeUseCase
	updateKnowledgeUseCase knowledge.UpdateKnowledgeUseCase
	deleteKnowledgeUseCase knowledge.DeleteKnowledgeUseCase
	searchKnowledgeUseCase knowledge.SearchKnowledgeUseCase
}

func NewKnowledgeHandler(
	createKnowledgeUseCase knowledge.CreateKnowledgeUseCase,
	updateKnowledgeUseCase knowledge.UpdateKnowledgeUseCase,
	deleteKnowledgeUseCase knowledge.DeleteKnowledgeUseCase,
	searchKnowledgeUseCase knowledge.SearchKnowledgeUseCase,
) *KnowledgeHandler {
	return &KnowledgeHandler{
		createKnowledgeUseCase: createKnowledgeUseCase,
		updateKnowledgeUseCase: updateKnowledgeUseCase,
		deleteKnowledgeUseCase: deleteKnowledgeUseCase,
		searchKnowledgeUseCase: searchKnowledgeUseCase,
	}
}

// CreateKnowledgeRequest represents the create knowledge request body
type CreateKnowledgeRequest struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Status  string   `json:"status"`
	TagIDs  []string `json:"tag_ids"`
}

// UpdateKnowledgeRequest represents the update knowledge request body
type UpdateKnowledgeRequest struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	TagIDs  []string `json:"tag_ids"`
}

// SearchKnowledgeRequest represents the search knowledge request query
type SearchKnowledgeRequest struct {
	Query    string   `query:"query"`
	TagIDs   []string `query:"tag_ids"`
	AuthorID string   `query:"author_id"`
}

// Create handles creating a new knowledge
// @Summary Create knowledge
// @Description Create a new knowledge
// @Tags knowledge
// @Accept json
// @Produce json
// @Param request body CreateKnowledgeRequest true "Knowledge data"
// @Security ApiKeyAuth
// @Success 201 {object} model.Knowledge
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge [post]
func (h *KnowledgeHandler) Create(c echo.Context) error {
	var req CreateKnowledgeRequest
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

	// Create knowledge
	knowledge, err := h.createKnowledgeUseCase.Execute(knowledge.CreateKnowledgeInput{
		Title:    req.Title,
		Content:  req.Content,
		Status:   req.Status,
		AuthorID: claims.UserID,
		TenantID: claims.TenantID,
		TagIDs:   req.TagIDs,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to create knowledge", err)
	}

	return appErrors.SendCreated(c, knowledge)
}

// Get handles getting a knowledge by ID
// @Summary Get knowledge
// @Description Get a knowledge by ID
// @Tags knowledge
// @Accept json
// @Produce json
// @Param id path string true "Knowledge ID"
// @Security ApiKeyAuth
// @Success 200 {object} model.Knowledge
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge/{id} [get]
func (h *KnowledgeHandler) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return appErrors.NewValidationError("ID is required", nil, nil)
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Search for knowledge with ID
	knowledges, err := h.searchKnowledgeUseCase.Execute(knowledge.SearchKnowledgeInput{
		Query:    id,
		TenantID: claims.TenantID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to get knowledge", err)
	}

	// Find the knowledge with the exact ID
	var foundKnowledge *model.Knowledge
	for _, k := range knowledges {
		if k.ID == id {
			foundKnowledge = k
			break
		}
	}

	if foundKnowledge == nil {
		return appErrors.NotFound("Knowledge not found", nil)
	}

	return appErrors.SendOK(c, foundKnowledge)
}

// Update handles updating a knowledge
// @Summary Update knowledge
// @Description Update a knowledge
// @Tags knowledge
// @Accept json
// @Produce json
// @Param id path string true "Knowledge ID"
// @Param request body UpdateKnowledgeRequest true "Knowledge data"
// @Security ApiKeyAuth
// @Success 200 {object} model.Knowledge
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge/{id} [put]
func (h *KnowledgeHandler) Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return appErrors.NewValidationError("ID is required", nil, nil)
	}

	var req UpdateKnowledgeRequest
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

	// Update knowledge
	knowledge, err := h.updateKnowledgeUseCase.Execute(knowledge.UpdateKnowledgeInput{
		ID:       id,
		Title:    req.Title,
		Content:  req.Content,
		TenantID: claims.TenantID,
		TagIDs:   req.TagIDs,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to update knowledge", err)
	}

	return appErrors.SendOK(c, knowledge)
}

// Delete handles deleting a knowledge
// @Summary Delete knowledge
// @Description Delete a knowledge
// @Tags knowledge
// @Accept json
// @Produce json
// @Param id path string true "Knowledge ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge/{id} [delete]
func (h *KnowledgeHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return appErrors.NewValidationError("ID is required", nil, nil)
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Delete knowledge
	err := h.deleteKnowledgeUseCase.Execute(knowledge.DeleteKnowledgeInput{
		ID:       id,
		TenantID: claims.TenantID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to delete knowledge", err)
	}

	return appErrors.SendNoContent(c)
}

// Search handles searching for knowledge
// @Summary Search knowledge
// @Description Search for knowledge
// @Tags knowledge
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param tag_ids query []string false "Tag IDs"
// @Param author_id query string false "Author ID"
// @Security ApiKeyAuth
// @Success 200 {array} model.Knowledge
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /knowledge [get]
func (h *KnowledgeHandler) Search(c echo.Context) error {
	var req SearchKnowledgeRequest
	if err := c.Bind(&req); err != nil {
		return appErrors.NewValidationError("Invalid request parameters", nil, err)
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Search knowledge
	knowledges, err := h.searchKnowledgeUseCase.Execute(knowledge.SearchKnowledgeInput{
		Query:    req.Query,
		TenantID: claims.TenantID,
		TagIDs:   req.TagIDs,
		AuthorID: req.AuthorID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to search knowledge", err)
	}

	return appErrors.SendOK(c, knowledges)
}

// RegisterRoutes registers the knowledge routes
func (h *KnowledgeHandler) RegisterRoutes(g *echo.Group) {
	knowledge := g.Group("/knowledge")
	knowledge.POST("", h.Create)
	knowledge.GET("", h.Search)
	knowledge.GET("/:id", h.Get)
	knowledge.PUT("/:id", h.Update)
	knowledge.DELETE("/:id", h.Delete)
}