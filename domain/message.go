package domain

import (
	"time"
)

type Message struct {
	ID        uint      `json:"id,omitempty"`
	Message   string    `json:"message,omitempty"`
	Star      bool      `json:"star,omitempty"`
	UserID    uint      `json:"user_id,omitempty"`
	ChatID    uint      `json:"chat_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
type Messages []*Message

type LastMessageFromPeople struct {
	ID        uint      `json:"id,omitempty"`
	Message   string    `json:"message,omitempty"`
	Star      bool      `json:"star,omitempty"`
	ChatID    uint      `json:"chat_id,omitempty"`
	UserID    uint      `json:"user_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Picture   string    `json:"picture,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
type LastMessageFromPeoples []*LastMessageFromPeople


type LastMessageFromGroup struct {
	ID          uint      `json:"id,omitempty"`
	Message     string    `json:"message,omitempty"`
	Star        bool      `json:"star,omitempty"`
	ChatID      uint      `json:"chat_id,omitempty"`
	UserID      uint      `json:"user_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
type LastMessageFromGroups []*LastMessageFromGroup


type LastMessageFromAll struct {
	ID          uint      `json:"id,omitempty"`
	Message     string    `json:"message,omitempty"`
	Star        bool      `json:"star,omitempty"`
	ChatID      uint      `json:"chat_id,omitempty"`
	UserID      uint      `json:"user_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Picture     string    `json:"picture,omitempty"`
	Group       bool           `json:"group,omitempty"`
	NameGroup   string    `json:"name_group,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
type LastMessageFromAlls []*LastMessageFromAll
