package persistence

import (
	"database/sql"
	"fmt"
	"github.com/MikelSot/melody/domain"
)

const (
	create_message=`INSERT INTO messages (message, user_id, chat_id) VALUES ($1, $2, $3) RETURNING id`
	update_star=`UPDATE messages SET star = true WHERE id= $1;`
	delete_message=`DELETE FROM messages WHERE id= $1`
)

type message struct {
	db *sql.DB
}

func NewMessage(db *sql.DB) *message {
	return &message{db}
}

func (m message) Create(d *domain.Message) error {
	stmt, err := m.db.Prepare(create_message)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(d.Message,d.UserID,d.ChatID).Scan(&d.ID)
	if err != nil {
		return err
	}
	return  nil
}

func (m message) UpdateStarField(id uint) error {
	stmt, err := m.db.Prepare(update_star)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se puedo actualizar el campo star" )
	}
	return nil
}

func (m message) DeleteMessage(id uint) error {
	stmt, err := m.db.Prepare(delete_message)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se puedo eliminar el mensaje" )
	}
	return nil
}
