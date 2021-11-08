package persistence

import (
	"database/sql"
	"fmt"
	"github.com/MikelSot/melody/domain"
)

const (
	create_chat=`INSERT INTO chats("group", name, description) VALUES ($1, $2, $3) RETURNING id`
	update_chat=`UPDATE chats SET name=$1, description=$2 WHERE id= $3`
	getByIdChat=`SELECT id, "group", name, description FROM chats WHERE id = $1`
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
	err = stmt.QueryRow(d.Group,d.Name,d.Description).Scan(&d.ID)
	if err != nil {
		return err
	}
	return  nil}

func (c chat) Update(d *domain.Chat) error {
	stmt, err := c.db.Prepare(update_chat)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(d.Name,d.Description,	d.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no existe un chat con id: %d", d.ID)
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
	err = stmt.QueryRow(id).Scan(
		ch.ID,
		ch.Group,
		ch.Name,
		ch.Description,
	)
	if err != nil {
		return domain.Chat{}, err
	}
	return ch, nil
}
