package persistence

import (
	"database/sql"
	"github.com/MikelSot/melody/domain"
)

const (
	create_chat=`INSERT INTO chats VALUES(DEFAULT) RETURNING id`
	getByIdChat=`SELECT id FROM chats WHERE id = $1`
)

type chat struct {
	db *sql.DB
}

func NewChat(db *sql.DB) *chat {
	return &chat{db}
}

func (c chat) Create(d *domain.Chat) error {
	stmt, err := c.db.Prepare(create_chat)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&d.ID)
	if err != nil {
		return err
	}
	return  nil
}

func (c chat) GetById(id uint) (domain.Chat, error) {
	stmt, err := c.db.Prepare(getByIdChat)
	if err != nil {
		return domain.Chat{}, err
	}
	defer stmt.Close()
	ch := domain.Chat{}
	err = stmt.QueryRow(id).Scan(ch.ID)
	if err != nil {
		return domain.Chat{}, err
	}
	return ch, nil
}
