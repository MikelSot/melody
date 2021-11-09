package domain

type ChatUser struct {
	ChatID uint `json:"chat_id,omitempty"`
	UserID uint `json:"user_id,omitempty"`
}

type  ChatUsers []*ChatUser

type GetChatFromPeople struct {
	ChatID uint `json:"chat_id,omitempty"`
	UserID uint `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
}

type GetChatFromPeoples []*GetChatFromPeople