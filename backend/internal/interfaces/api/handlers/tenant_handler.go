package handlers

import (
	"github.com/labstack/echo/v4"

	appErrors "github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/errors"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/tenant"
)

type TenantHandler struct {
	createTenantUseCase        tenant.CreateTenantUseCase
	updateTenantSettingsUseCase tenant.UpdateTenantSettingsUseCase
	deleteTenantUseCase        tenant.DeleteTenantUseCase
}

func NewTenantHandler(
	createTenantUseCase tenant.CreateTenantUseCase,
	updateTenantSettingsUseCase tenant.UpdateTenantSettingsUseCase,
	deleteTenantUseCase tenant.DeleteTenantUseCase,
) *TenantHandler {
	return &TenantHandler{
		createTenantUseCase:        createTenantUseCase,
		updateTenantSettingsUseCase: updateTenantSettingsUseCase,
		deleteTenantUseCase:        deleteTenantUseCase,
	}
}

// CreateTenantRequest represents the create tenant request body
type CreateTenantRequest struct {
	Name   string `json:"name" validate:"required"`
	Domain string `json:"domain" validate:"required"`
	Settings struct {
		Theme struct {
			PrimaryColor   string `json:"primary_color" validate:"required,hexcolor"`
			SecondaryColor string `json:"secondary_color" validate:"required,hexcolor"`
		} `json:"theme" validate:"required"`
		Features struct {
			Comments bool `json:"comments"`
			Tags     bool `json:"tags"`
			Ratings  bool `json:"ratings"`
		} `json:"features" validate:"required"`
	} `json:"settings" validate:"required"`
}

// UpdateTenantSettingsRequest represents the update tenant settings request body
type UpdateTenantSettingsRequest struct {
	Settings struct {
		Theme struct {
			PrimaryColor   string `json:"primary_color" validate:"required,hexcolor"`
			SecondaryColor string `json:"secondary_color" validate:"required,hexcolor"`
		} `json:"theme" validate:"required"`
		Features struct {
			Comments bool `json:"comments"`
			Tags     bool `json:"tags"`
			Ratings  bool `json:"ratings"`
		} `json:"features" validate:"required"`
	} `json:"settings" validate:"required"`
}

// Create handles creating a new tenant
// @Summary Create tenant
// @Description Create a new tenant
// @Tags tenants
// @Accept json
// @Produce json
// @Param request body CreateTenantRequest true "Tenant data"
// @Security ApiKeyAuth
// @Success 201 {object} model.Tenant
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tenants [post]
func (h *TenantHandler) Create(c echo.Context) error {
	var req CreateTenantRequest
	if err := c.Bind(&req); err != nil {
		return appErrors.NewValidationError("Invalid request body", nil, err)
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	// Create tenant
	tenant, err := h.createTenantUseCase.Execute(tenant.CreateTenantInput{
		Name:   req.Name,
		Domain: req.Domain,
		Theme: model.Theme{
			PrimaryColor:   req.Settings.Theme.PrimaryColor,
			SecondaryColor: req.Settings.Theme.SecondaryColor,
		},
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to create tenant", err)
	}

	return appErrors.SendCreated(c, tenant)
}

// UpdateSettings handles updating tenant settings
// @Summary Update tenant settings
// @Description Update tenant settings
// @Tags tenants
// @Accept json
// @Produce json
// @Param id path string true "Tenant ID"
// @Param request body UpdateTenantSettingsRequest true "Tenant settings data"
// @Security ApiKeyAuth
// @Success 200 {object} model.Tenant
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 403 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tenants/{id}/settings [put]
func (h *TenantHandler) UpdateSettings(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return appErrors.NewValidationError("ID is required", nil, nil)
	}

	var req UpdateTenantSettingsRequest
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

	// Check if user is admin
	if claims.Role != "admin" {
		return appErrors.Forbidden("Only admins can update tenant settings", nil)
	}

	// Check if user belongs to the tenant
	if claims.TenantID != id {
		return appErrors.Forbidden("You can only update your own tenant", nil)
	}

	// Update tenant settings
	tenant, err := h.updateTenantSettingsUseCase.Execute(tenant.UpdateTenantSettingsInput{
		ID: id,
		Settings: model.Settings{
			Theme: model.Theme{
				PrimaryColor:   req.Settings.Theme.PrimaryColor,
				SecondaryColor: req.Settings.Theme.SecondaryColor,
			},
			Features: model.Features{
				Comments: req.Settings.Features.Comments,
				Tags:     req.Settings.Features.Tags,
				Ratings:  req.Settings.Features.Ratings,
			},
		},
	})
	if err != nil {
		return appErrors.InternalServerError("Failed to update tenant settings", err)
	}

	return appErrors.SendOK(c, tenant)
}

// Delete handles deleting a tenant
// @Summary Delete tenant
// @Description Delete a tenant
// @Tags tenants
// @Accept json
// @Produce json
// @Param id path string true "Tenant ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 403 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tenants/{id} [delete]
func (h *TenantHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return appErrors.NewValidationError("ID is required", nil, nil)
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Check if user is admin
	if claims.Role != "admin" {
		return appErrors.Forbidden("Only admins can delete tenants", nil)
	}

	// Check if user belongs to the tenant
	if claims.TenantID != id {
		return appErrors.Forbidden("You can only delete your own tenant", nil)
	}

	// Delete tenant
	err := h.deleteTenantUseCase.Execute(id)
	if err != nil {
		return appErrors.InternalServerError("Failed to delete tenant", err)
	}

	return appErrors.SendNoContent(c)
}

// Get handles getting a tenant
// @Summary Get tenant
// @Description Get a tenant by ID
// @Tags tenants
// @Accept json
// @Produce json
// @Param id path string true "Tenant ID"
// @Security ApiKeyAuth
// @Success 200 {object} model.Tenant
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 401 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tenants/{id} [get]
func (h *TenantHandler) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return appErrors.NewValidationError("ID is required", nil, nil)
	}

	// Get user claims from context
	claims := getUserClaims(c)
	if claims == nil {
		return appErrors.Unauthorized("Authentication required", nil)
	}

	// Get repositories
	repo := c.Get("repositories").(RepositoriesProvider).Tenant()

	// Get tenant
	tenant, err := repo.FindByID(id)
	if err != nil {
		return appErrors.NotFound("Tenant not found", err)
	}

	return appErrors.SendOK(c, tenant)
}

// GetByDomain handles getting a tenant by domain
// @Summary Get tenant by domain
// @Description Get a tenant by domain
// @Tags tenants
// @Accept json
// @Produce json
// @Param domain path string true "Tenant domain"
// @Success 200 {object} model.Tenant
// @Failure 400 {object} appErrors.ErrorResponse
// @Failure 404 {object} appErrors.ErrorResponse
// @Failure 500 {object} appErrors.ErrorResponse
// @Router /tenants/domain/{domain} [get]
func (h *TenantHandler) GetByDomain(c echo.Context) error {
	domain := c.Param("domain")
	if domain == "" {
		return appErrors.NewValidationError("Domain is required", nil, nil)
	}

	// Get repositories
	repo := c.Get("repositories").(RepositoriesProvider).Tenant()

	// Get tenant
	tenant, err := repo.FindByDomain(domain)
	if err != nil {
		return appErrors.NotFound("Tenant not found", err)
	}

	return appErrors.SendOK(c, tenant)
}

// RegisterRoutes registers the tenant routes
func (h *TenantHandler) RegisterRoutes(g *echo.Group) {
	tenants := g.Group("/tenants")
	tenants.POST("", h.Create)
	tenants.GET("/:id", h.Get)
	tenants.GET("/domain/:domain", h.GetByDomain)
	tenants.PUT("/:id/settings", h.UpdateSettings)
	tenants.DELETE("/:id", h.Delete)
}