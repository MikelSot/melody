package query

import (
	"database/sql"
	"github.com/MikelSot/melody/domain"
)

const (
	getAllStarred=`select id, message, star, user_id, chat_id from messages WHERE star = true`

	getMessageFromChat=`select id, message,star, user_id, chat_id, created_at from messages
						where id  in (select id from messages where chat_id = $1
									  order by id desc limit $2)`
	)

type message struct {
	db *sql.DB
}

func NewMessage(db *sql.DB) *message {
	return &message{db}
}

func (m message) GetMessageFromChat(max, myId uint) (domain.Messages, error) {
	stmt, err := m.db.Prepare(getMessageFromChat)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(myId, max)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make(domain.Messages, 0)
	for rows.Next() {
		ms := &domain.Message{}
		err := rows.Scan(
			&ms.ID,
			&ms.Message,
			&ms.Star,
			&ms.UserID,
			&ms.ChatID,
			&ms.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, ms)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (m message) GetAllStarred() (domain.Messages, error) {
	stmt, err := m.db.Prepare(getAllStarred)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make(domain.Messages, 0)
	for rows.Next() {
		ms := &domain.Message{}
		err := rows.Scan(
			&ms.ID,
			&ms.Message,
			&ms.Star,
			&ms.UserID,
			&ms.ChatID,
			&ms.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, ms)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
