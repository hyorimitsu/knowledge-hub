package tenant

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type createTenantUseCase struct {
	tenantRepository repository.TenantRepository
}

// NewCreateTenantUseCase creates a new instance of CreateTenantUseCase
func NewCreateTenantUseCase(tenantRepository repository.TenantRepository) CreateTenantUseCase {
	return &createTenantUseCase{
		tenantRepository: tenantRepository,
	}
}

// Execute creates a new tenant
func (uc *createTenantUseCase) Execute(input CreateTenantInput) (*model.Tenant, error) {
	// Validate input
	if input.Name == "" {
		return nil, errors.New("tenant name is required")
	}
	if input.Domain == "" {
		return nil, errors.New("tenant domain is required")
	}

	// Check if domain already exists
	existingTenant, err := uc.tenantRepository.FindByDomain(input.Domain)
	if err == nil && existingTenant != nil {
		return nil, errors.New("tenant domain already exists")
	}

	// Create tenant
	now := time.Now()
	tenant := &model.Tenant{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Domain:    input.Domain,
		Settings: model.Settings{
			Theme: input.Theme,
			Features: model.Features{
				Comments: true,
				Tags:     true,
				Ratings:  true,
			},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Save tenant
	err = uc.tenantRepository.Create(tenant)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}