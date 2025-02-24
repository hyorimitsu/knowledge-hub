package persistence

import (
	"gorm.io/gorm"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type knowledgeRepository struct {
	db *Database
}

func NewKnowledgeRepository(db *Database) repository.KnowledgeRepository {
	return &knowledgeRepository{db}
}

func (r *knowledgeRepository) Create(knowledge *model.Knowledge) error {
	return r.db.Create(knowledge).Error
}

func (r *knowledgeRepository) FindByID(id string, tenantID string) (*model.Knowledge, error) {
	var knowledge model.Knowledge
	err := r.db.
		Preload("Tags").
		Preload("Comments").
		First(&knowledge, "id = ? AND tenant_id = ?", id, tenantID).
		Error
	if err != nil {
		return nil, err
	}
	return &knowledge, nil
}

func (r *knowledgeRepository) FindAll(tenantID string) ([]*model.Knowledge, error) {
	var knowledges []*model.Knowledge
	err := r.db.
		Preload("Tags").
		Preload("Comments").
		Where("tenant_id = ?", tenantID).
		Find(&knowledges).
		Error
	if err != nil {
		return nil, err
	}
	return knowledges, nil
}

func (r *knowledgeRepository) Update(knowledge *model.Knowledge) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Update tags
		if err := tx.Model(knowledge).Association("Tags").Replace(knowledge.Tags); err != nil {
			return err
		}

		// Update knowledge
		return tx.Save(knowledge).Error
	})
}

func (r *knowledgeRepository) Delete(id string, tenantID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Delete related comments
		if err := tx.Where("knowledge_id = ? AND tenant_id = ?", id, tenantID).Delete(&model.Comment{}).Error; err != nil {
			return err
		}

		// Delete knowledge_tags associations
		knowledge := &model.Knowledge{ID: id}
		if err := tx.Model(knowledge).Association("Tags").Clear(); err != nil {
			return err
		}

		// Delete knowledge
		return tx.Delete(&model.Knowledge{}, "id = ? AND tenant_id = ?", id, tenantID).Error
	})
}
