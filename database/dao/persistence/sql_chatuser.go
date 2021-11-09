package persistence

import (
	"database/sql"
	"github.com/MikelSot/melody/domain"
)

const (
	create_chatuser=`INSERT INTO chat_users (chat_id, user_id) VALUES ($1, $2)`
)

type chatUser struct {
	db *sql.DB
}

func NewChatUser(db *sql.DB) *chatUser {
	return &chatUser{db}
}

func (c chatUser) Create(cu *domain.ChatUser) error {
	stmt, err := c.db.Prepare(create_chatuser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&cu.ChatID,&cu.UserID)
	if err != nil {
		return err
	}
	return  nil
}
