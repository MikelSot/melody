package query

import (
	"database/sql"
	"github.com/MikelSot/melody/domain"
)

const (
	match =`select chat_id, user_id from chat_users
			where chat_id in (select chat_id from chat_users
							  inner join chats c on c.id = chat_users.chat_id
                  			  where c."group" <> true and user_id = $2) and  user_id = $1;`

	allPeopleAddedToMyChat=`select chat_id, u.id, u.name, u.last_name from chat_users
							inner join users u on u.id = chat_users.user_id
							where chat_id in(select chat_id from chat_users
                    						 inner join chats c on c.id = chat_users.chat_id
                    						 where user_id =$2 and c."group" = false) and user_id <> $2 
							limit $1;`
)

type chatUser struct {
	db *sql.DB
}

func NewChatUser(db *sql.DB) *chatUser {
	return &chatUser{db}
}

func (c chatUser) Match(myId uint, yourId uint) (domain.ChatUser, error) {
	stmt, err := c.db.Prepare(match)
	if err != nil {
		return domain.ChatUser{}, err
	}
	defer stmt.Close()
	cu := domain.ChatUser{}
	stmt.QueryRow(myId, yourId).Scan(cu.ChatID, cu.UserID)
	if err != nil {
		return domain.ChatUser{}, err
	}
	return cu, nil
}

func (c chatUser) AllPeopleAddedToMyChat(max uint, myId uint) (domain.GetChatFromPeoples, error) {
	stmt, err := c.db.Prepare(allPeopleAddedToMyChat)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(max, myId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	chats := make(domain.GetChatFromPeoples, 0)
	for rows.Next() {
		chat := &domain.GetChatFromPeople{}
		err := rows.Scan(&chat.ChatID,&chat.UserID, &chat.Name, &chat.LastName)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chats, nil
}