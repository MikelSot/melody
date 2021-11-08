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

	getLatestChatMessageFromPeopleAdded=`select messages.id,message, star, chat_id, user_id,u.name,u.last_name, 
										 u.picture, messages.created_at from messages
										 inner join users u on u.id = messages.user_id
										 where chat_id =$1 order by id desc
										 limit 1`

	getLatestChatMessageFromGroupAdded=`select messages.id, message,star,chat_id,user_id,c.name,
										messages.created_at from messages
										inner join chats c on c.id = messages.chat_id
										where chat_id = $1 order by id desc
										limit 1`

	getLatestChatMessageFromAllAdded=`select messages.id, message, star, chat_id, user_id, u.name, u.last_name,
									  u.picture, c."group", c.name, messages.created_at  from messages
									  inner join chats c on c.id = messages.chat_id
									  inner join users u on u.id = messages.user_id
									  where chat_id = $1 order by id desc
									  limit 1`
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

func (m message) GetLatestChatMessageFromPeopleAdded(chatID uint) (domain.LastMessageFromPeoples, error) {
	stmt, err := m.db.Prepare(getLatestChatMessageFromPeopleAdded)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make(domain.LastMessageFromPeoples, 0)
	for rows.Next() {
		ms := &domain.LastMessageFromPeople{}
		err := rows.Scan(
			&ms.ID,
			&ms.Message,
			&ms.Star,
			&ms.ChatID,
			&ms.UserID,
			&ms.Name,
			&ms.LastName,
			&ms.Picture,
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

func (m message) GetLatestChatMessageFromGroupAdded(chatID uint) (domain.LastMessageFromGroups, error) {
	stmt, err := m.db.Prepare(getLatestChatMessageFromGroupAdded)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make(domain.LastMessageFromGroups, 0)
	for rows.Next() {
		ms := &domain.LastMessageFromGroup{}
		err := rows.Scan(
			&ms.ID,
			&ms.Message,
			&ms.Star,
			&ms.ChatID,
			&ms.UserID,
			&ms.Name,
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

func (m message) GetLatestChatMessageFromAllAdded(chatID uint) (domain.LastMessageFromAlls, error) {
	stmt, err := m.db.Prepare(getLatestChatMessageFromAllAdded)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make(domain.LastMessageFromAlls, 0)
	for rows.Next() {
		ms := &domain.LastMessageFromAll{}
		err := rows.Scan(
			&ms.ID,
			&ms.Message,
			&ms.Star,
			&ms.ChatID,
			&ms.UserID,
			&ms.Name,
			&ms.LastName,
			&ms.Picture,
			&ms.Group,
			&ms.NameGroup,
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

