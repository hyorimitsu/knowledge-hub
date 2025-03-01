package tenant

import (
	"errors"
	"time"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type updateTenantSettingsUseCase struct {
	tenantRepository repository.TenantRepository
}

// NewUpdateTenantSettingsUseCase creates a new instance of UpdateTenantSettingsUseCase
func NewUpdateTenantSettingsUseCase(tenantRepository repository.TenantRepository) UpdateTenantSettingsUseCase {
	return &updateTenantSettingsUseCase{
		tenantRepository: tenantRepository,
	}
}

// Execute updates tenant settings
func (uc *updateTenantSettingsUseCase) Execute(input UpdateTenantSettingsInput) (*model.Tenant, error) {
	// Validate input
	if input.ID == "" {
		return nil, errors.New("tenant ID is required")
	}

	// Find tenant
	tenant, err := uc.tenantRepository.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, errors.New("tenant not found")
	}

	// Update settings
	tenant.Settings = input.Settings
	tenant.UpdatedAt = time.Now()

	// Save tenant
	err = uc.tenantRepository.Update(tenant)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}