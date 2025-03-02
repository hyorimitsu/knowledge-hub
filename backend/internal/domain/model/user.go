package model

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-" gorm:"column:password_hash"` // Password is not exposed in JSON
	Avatar    string    `json:"avatar,omitempty" gorm:"-"` // Skip this field when inserting into the database
	TenantID  string    `json:"tenant_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleEditor Role = "editor"
	RoleViewer Role = "viewer"
)

func (r Role) String() string {
	return string(r)
}

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleEditor, RoleViewer:
		return true
	default:
		return false
	}
}
