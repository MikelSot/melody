package domain

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID            uint           `gorm:"primaryKey" json:"id,omitempty"`
	Message       string         `gorm:"type:varchar(255); default:''; not null" json:"message,omitempty"`
	Archive       string         `gorm:"type:varchar(200)" json:"archive,omitempty"`
	Star          bool           `gorm:"default:false" json:"star,omitempty"`
	UserID        uint           `gorm:"not null" json:"user_id,omitempty"`
	ChatID        uint           `gorm:"not null" json:"chat_id,omitempty"`
	TypeMessageID uint8          `gorm:"not null" json:"type_message_id,omitempty"`
	CreatedAt     time.Time      `gorm:"default:now()" json:"created_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Messages []Message