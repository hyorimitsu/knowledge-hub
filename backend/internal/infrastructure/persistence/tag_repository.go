package persistence

import (
	"gorm.io/gorm"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type tagRepository struct {
	db *Database
}

func NewTagRepository(db *Database) repository.TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) Create(tag *model.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) FindByID(id string, tenantID string) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.First(&tag, "id = ? AND tenant_id = ?", id, tenantID).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) FindAll(tenantID string) ([]*model.Tag, error) {
	var tags []*model.Tag
	err := r.db.Where("tenant_id = ?", tenantID).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) Update(tag *model.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepository) Delete(id string, tenantID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Get the tag to delete
		tag := &model.Tag{ID: id}

		// Remove associations with knowledge
		if err := tx.Model(tag).Association("Knowledge").Clear(); err != nil {
			return err
		}

		// Delete the tag
		return tx.Delete(&model.Tag{}, "id = ? AND tenant_id = ?", id, tenantID).Error
	})
}
