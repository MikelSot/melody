package domain

import (
	"time"
)

type Chat struct {
	ID          uint           `json:"id,omitempty"`
	Group       bool           `json:"group,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
}
