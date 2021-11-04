package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id,omitempty"`
	Name        string         `gorm:"type:varchar(80); default:''; not null" json:"name,omitempty"`
	LastName    string         `gorm:"type:varchar(80); default:''; not null" json:"last_name,omitempty"`
	Nickname    string         `gorm:"type:varchar(80); default:''" json:"nickname,omitempty"`
	Email       string         `gorm:"uniqueIndex; type:varchar(150); not null" json:"email,omitempty"`
	Password    string         `gorm:"type:varchar(200); not null" json:"password,omitempty"`
	Online      bool           `gorm:"default:false" json:"online,omitempty"`
	Description string         `gorm:"type:varchar(150); default:''" json:"description,omitempty"`
	Picture     string         `gorm:"type:varchar(200)" json:"picture,omitempty"`
	CreatedAt   time.Time      `gorm:"default:now()" json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Messages    []Message
}
