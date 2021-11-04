package domain

import (
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	ID          uint           `gorm:"primaryKey" json:"id,omitempty"`
	Group       bool           `gorm:"default:false" json:"group,omitempty"`
	Name        string         `gorm:"type:varchar(100); default:''; not null" json:"name,omitempty"`
	Description string         `gorm:"type:varchar(150); default:''" json:"description,omitempty"`
	CreatedAt   time.Time      `gorm:"default:now()" json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Users       []User         `gorm:"many2many:chat_users;"`
	Messages    []Message
}
