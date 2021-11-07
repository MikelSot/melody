package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint           `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	LastName    string         `json:"last_name,omitempty"`
	Nickname    string         `json:"nickname,omitempty"`
	Email       string         `json:"email,omitempty"`
	Password    string         `json:"password,omitempty"`
	Online      bool           `json:"online,omitempty"`
	Description string         `json:"description,omitempty"`
	Picture     string         `json:"picture,omitempty"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}
type Users []User

type Register struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Nickname struct {
	Nickname string `json:"nickname,omitempty"`
}
