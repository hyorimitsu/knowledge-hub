package persistence

import (
	"gorm.io/gorm"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type userRepository struct {
	db *Database
}

func NewUserRepository(db *Database) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id string, tenantID string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "id = ? AND tenant_id = ?", id, tenantID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string, tenantID string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ? AND tenant_id = ?", email, tenantID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id string, tenantID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Update related records to set author_id to null
		if err := tx.Model(&model.Knowledge{}).
			Where("author_id = ? AND tenant_id = ?", id, tenantID).
			Update("author_id", nil).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Comment{}).
			Where("author_id = ? AND tenant_id = ?", id, tenantID).
			Update("author_id", nil).Error; err != nil {
			return err
		}

		// Delete user
		return tx.Delete(&model.User{}, "id = ? AND tenant_id = ?", id, tenantID).Error
	})
}
