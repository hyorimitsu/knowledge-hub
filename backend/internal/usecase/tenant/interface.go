package tenant

import "github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"

// CreateTenantUseCase defines the interface for creating a tenant
type CreateTenantUseCase interface {
	Execute(input CreateTenantInput) (*model.Tenant, error)
}

// CreateTenantInput contains the data needed to create a tenant
type CreateTenantInput struct {
	Name   string
	Domain string
	Theme  model.Theme
}

// UpdateTenantSettingsUseCase defines the interface for updating tenant settings
type UpdateTenantSettingsUseCase interface {
	Execute(input UpdateTenantSettingsInput) (*model.Tenant, error)
}

// UpdateTenantSettingsInput contains the data needed to update tenant settings
type UpdateTenantSettingsInput struct {
	ID       string
	Settings model.Settings
}

// DeleteTenantUseCase defines the interface for deleting a tenant
type DeleteTenantUseCase interface {
	Execute(id string) error
}