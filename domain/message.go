package domain

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID            uint           `json:"id,omitempty"`
	Message       string         `json:"message,omitempty"`
	Archive       string         `json:"archive,omitempty"`
	Star          bool           `json:"star,omitempty"`
	UserID        uint           `json:"user_id,omitempty"`
	ChatID        uint           `json:"chat_id,omitempty"`
	TypeMessageID uint8          `json:"type_message_id,omitempty"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type Messages []Message