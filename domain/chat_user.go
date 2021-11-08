package domain

type ChatUser struct {
	ChatID uint `json:"chat_id,omitempty"`
	UserID uint `json:"user_id,omitempty"`
}

type  ChatUsers []*ChatUser