package model

import "time"

type Knowledge struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"author_id"`
	TenantID  string    `json:"tenant_id"`
	Tags      []Tag     `json:"tags" gorm:"many2many:knowledge_tags;"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:KnowledgeID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tag struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	TenantID  string    `json:"tenant_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Content     string    `json:"content"`
	AuthorID    string    `json:"author_id"`
	KnowledgeID string    `json:"knowledge_id"`
	TenantID    string    `json:"tenant_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
