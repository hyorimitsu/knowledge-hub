package persistence

import (
	"gorm.io/gorm"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type tenantRepository struct {
	db *Database
}

func NewTenantRepository(db *Database) repository.TenantRepository {
	return &tenantRepository{db}
}

func (r *tenantRepository) Create(tenant *model.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *tenantRepository) FindByID(id string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.First(&tenant, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) FindByDomain(domain string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.First(&tenant, "domain = ?", domain).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) Update(tenant *model.Tenant) error {
	return r.db.Save(tenant).Error
}

func (r *tenantRepository) Delete(id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Delete related records first
		if err := tx.Where("tenant_id = ?", id).Delete(&model.User{}).Error; err != nil {
			return err
		}
		if err := tx.Where("tenant_id = ?", id).Delete(&model.Knowledge{}).Error; err != nil {
			return err
		}
		if err := tx.Where("tenant_id = ?", id).Delete(&model.Tag{}).Error; err != nil {
			return err
		}
		if err := tx.Where("tenant_id = ?", id).Delete(&model.Comment{}).Error; err != nil {
			return err
		}

		// Delete tenant
		return tx.Delete(&model.Tenant{}, "id = ?", id).Error
	})
}
