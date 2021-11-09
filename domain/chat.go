package domain

import (
	"time"
)

type Chat struct {
	ID          uint           `json:"id,omitempty"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
}
