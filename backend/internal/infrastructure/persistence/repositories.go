package persistence

import (
	"gorm.io/gorm"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/domain/repository"
)

type Repositories struct {
	db        *gorm.DB
	tenant    repository.TenantRepository
	user      repository.UserRepository
	knowledge repository.KnowledgeRepository
	tag       repository.TagRepository
	comment   repository.CommentRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		db:        db,
		tenant:    NewTenantRepository(&Database{db}),
		user:      NewUserRepository(&Database{db}),
		knowledge: NewKnowledgeRepository(&Database{db}),
		tag:       NewTagRepository(&Database{db}),
		comment:   NewCommentRepository(&Database{db}),
	}
}

func (r *Repositories) Tenant() repository.TenantRepository {
	return r.tenant
}

func (r *Repositories) User() repository.UserRepository {
	return r.user
}

func (r *Repositories) Knowledge() repository.KnowledgeRepository {
	return r.knowledge
}

func (r *Repositories) Tag() repository.TagRepository {
	return r.tag
}

func (r *Repositories) Comment() repository.CommentRepository {
	return r.comment
}

func (r *Repositories) DB() *gorm.DB {
	return r.db
}
