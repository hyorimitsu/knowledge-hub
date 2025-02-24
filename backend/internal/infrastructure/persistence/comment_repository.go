package persistence

import (
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type commentRepository struct {
	db *Database
}

func NewCommentRepository(db *Database) repository.CommentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) FindByID(id string, tenantID string) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.First(&comment, "id = ? AND tenant_id = ?", id, tenantID).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) FindByKnowledgeID(knowledgeID string, tenantID string) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := r.db.
		Where("knowledge_id = ? AND tenant_id = ?", knowledgeID, tenantID).
		Order("created_at DESC").
		Find(&comments).
		Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) Update(comment *model.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) Delete(id string, tenantID string) error {
	return r.db.Delete(&model.Comment{}, "id = ? AND tenant_id = ?", id, tenantID).Error
}
