package domain

import (
	"time"
)

type User struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Nickname    string    `json:"nickname,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	Online      bool      `json:"online,omitempty"`
	Description string    `json:"description,omitempty"`
	Picture     string    `json:"picture,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
type Users []User

type CreateUser struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UpdateUser struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Description string `json:"description,omitempty"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type FieldString struct {
	ID    uint   `json:"id,omitempty"`
	Field string `json:"Field,omitempty"`
}

type State struct {
	ID    uint `json:"id,omitempty"`
	Field bool `json:"state,omitempty"`
}
