package repository

import "github.com/hyorimitsu/knowledge-hub/backend/internal/domain/model"

type TenantRepository interface {
	Create(tenant *model.Tenant) error
	FindByID(id string) (*model.Tenant, error)
	FindByDomain(domain string) (*model.Tenant, error)
	Update(tenant *model.Tenant) error
	Delete(id string) error
}

type KnowledgeRepository interface {
	Create(knowledge *model.Knowledge) error
	FindByID(id string, tenantID string) (*model.Knowledge, error)
	FindAll(tenantID string) ([]*model.Knowledge, error)
	Search(query string, tenantID string, tagIDs []string, authorID string) ([]*model.Knowledge, error)
	Update(knowledge *model.Knowledge) error
	Delete(id string, tenantID string) error
}

type TagRepository interface {
	Create(tag *model.Tag) error
	FindByID(id string, tenantID string) (*model.Tag, error)
	FindAll(tenantID string) ([]*model.Tag, error)
	Update(tag *model.Tag) error
	Delete(id string, tenantID string) error
}

type CommentRepository interface {
	Create(comment *model.Comment) error
	FindByID(id string, tenantID string) (*model.Comment, error)
	FindByKnowledgeID(knowledgeID string, tenantID string) ([]*model.Comment, error)
	Update(comment *model.Comment) error
	Delete(id string, tenantID string) error
}

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id string, tenantID string) (*model.User, error)
	FindByEmail(email string, tenantID string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string, tenantID string) error
}
