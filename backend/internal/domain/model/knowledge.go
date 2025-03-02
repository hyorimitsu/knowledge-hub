package model

import "time"

// Knowledge status constants
const (
	KnowledgeStatusDraft     = "draft"
	KnowledgeStatusPublished = "published"
	KnowledgeStatusArchived  = "archived"
)

type Knowledge struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"author_id"`
	TenantID  string    `json:"tenant_id"`
	Status    string    `json:"status"`
	Tags      []Tag     `json:"tags" gorm:"many2many:knowledge_tags;"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:KnowledgeID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Knowledge
func (Knowledge) TableName() string {
	return "knowledge"
}

type Tag struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	TenantID  string    `json:"tenant_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Tag
func (Tag) TableName() string {
	return "tags"
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

// TableName specifies the table name for Comment
func (Comment) TableName() string {
	return "comments"
}
