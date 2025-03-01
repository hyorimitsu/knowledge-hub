package tenant

import (
	"errors"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type deleteTenantUseCase struct {
	tenantRepository repository.TenantRepository
}

// NewDeleteTenantUseCase creates a new instance of DeleteTenantUseCase
func NewDeleteTenantUseCase(tenantRepository repository.TenantRepository) DeleteTenantUseCase {
	return &deleteTenantUseCase{
		tenantRepository: tenantRepository,
	}
}

// Execute deletes a tenant
func (uc *deleteTenantUseCase) Execute(id string) error {
	// Validate input
	if id == "" {
		return errors.New("tenant ID is required")
	}

	// Find tenant
	tenant, err := uc.tenantRepository.FindByID(id)
	if err != nil {
		return err
	}
	if tenant == nil {
		return errors.New("tenant not found")
	}

	// Delete tenant
	return uc.tenantRepository.Delete(id)
}