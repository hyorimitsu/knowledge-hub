package handlers

import (
	"github.com/labstack/echo/v4"

	appErrors "github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/errors"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/tag"
)

type TagHandler struct {
	createTagUseCase tag.CreateTagUseCase
	updateTagUseCase tag.UpdateTagUseCase
	deleteTagUseCase tag.DeleteTagUseCase
}

func NewTagHandler(
	createTagUseCase tag.CreateTagUseCase,
	updateTagUseCase tag.UpdateTagUseCase,
	deleteTagUseCase tag.DeleteTagUseCase,
) *TagHandler {
	return &TagHandler{
		createTagUseCase: createTagUseCase,
		updateTagUseCase: updateTagUseCase,
		deleteTagUseCase: deleteTagUseCase,
	}
}

// CreateTagRequest represents the create tag request body
type CreateTagRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateTagRequest represents the update tag request body
type UpdateTagRequest struct {
	Name string `json:"name" validate:"required"`
}

// Create handles creating a new tag
// @Summary Create tag
// @Description Create a new tag
// @Tags tags
// @Accept json
// @Produce json
// @Param request body CreateTagRequest true "Tag data"
// @Security ApiKeyAuth
// @Success 201 {object} model.Tag
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tags [post]
func (h *TagHandler) Create(c echo.Context) error {
	var req CreateTagRequest
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

	// Create tag
	tag, err := h.createTagUseCase.Execute(tag.CreateTagInput{
		Name:     req.Name,
		TenantID: claims.TenantID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to create tag", err)
	}

	return appErrors.SendCreated(c, tag)
}

// Update handles updating a tag
// @Summary Update tag
// @Description Update a tag
// @Tags tags
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Param request body UpdateTagRequest true "Tag data"
// @Security ApiKeyAuth
// @Success 200 {object} model.Tag
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tags/{id} [put]
func (h *TagHandler) Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return appErrors.NewValidationError("ID is required", nil, nil)
	}

	var req UpdateTagRequest
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

	// Update tag
	tag, err := h.updateTagUseCase.Execute(tag.UpdateTagInput{
		ID:       id,
		Name:     req.Name,
		TenantID: claims.TenantID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to update tag", err)
	}

	return appErrors.SendOK(c, tag)
}

// Delete handles deleting a tag
// @Summary Delete tag
// @Description Delete a tag
// @Tags tags
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tags/{id} [delete]
func (h *TagHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return appErrors.NewValidationError("ID is required", nil, nil)
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Delete tag
	err := h.deleteTagUseCase.Execute(tag.DeleteTagInput{
		ID:       id,
		TenantID: claims.TenantID,
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to delete tag", err)
	}

	return appErrors.SendNoContent(c)
}

// List handles listing all tags
// @Summary List tags
// @Description List all tags
// @Tags tags
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} model.Tag
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tags [get]
func (h *TagHandler) List(c echo.Context) error {
	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Get tags repository
	repo := c.Get("repositories").(RepositoriesProvider).Tag()

	// List tags
	tags, err := repo.FindAll(claims.TenantID)
	if err != nil {
		return appErrors.InternalServerError("Failed to list tags", err)
	}

	return appErrors.SendOK(c, tags)
}

// RegisterRoutes registers the tag routes
func (h *TagHandler) RegisterRoutes(g *echo.Group) {
	tags := g.Group("/tags")
	tags.POST("", h.Create)
	tags.GET("", h.List)
	tags.PUT("/:id", h.Update)
	tags.DELETE("/:id", h.Delete)
}