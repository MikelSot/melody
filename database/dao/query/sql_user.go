package query

import (
	"database/sql"
	"github.com/MikelSot/melody/domain"
)

const (
	getAllNotAddedUsers = `select id, name,last_name, nick_name, online, picture from users
 						   where id not in (select user_id from chat_users
											where chat_id in(select chat_id from chat_users
                                         					 inner join chats c on c.id = chat_users.chat_id
                                         					 where user_id =$1 and c."group" = false)
											and user_id <> $1) and id <> $1
						   limit $2`
	getByIdUser    = `Select id, name, last_name, nick_name, email, description, picture from users where id = $1`
	emailExists    = `Select id, name, last_name, nick_name, email, description, picture from users where email = $1`
	nicknameExists = `Select id, name, last_name, nick_name, email, description, picture from users where nick_name = $1`
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *user {
	return &user{db}
}

func (u user) GetAllNotAddedUsers(max int, myid uint) (domain.Users, error) {
	stmt, err := u.db.Prepare(getAllNotAddedUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(myid, max)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make(domain.Users, 0)
	for rows.Next() {
		us := &domain.User{}
		err := rows.Scan(
			&us.ID,
			&us.Name,
			&us.Nickname,
			&us.Online,
			&us.Picture,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, us)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u user) GetById(id uint) (domain.User, error) {
	stmt, err := u.db.Prepare(getByIdUser)
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()
	us, err := u.scanData(id, stmt)
	if err != nil {
		return domain.User{}, err
	}
	return us, nil
}

func (u user) EmailFieldExists(email string) (domain.User, error) {
	stmt, err := u.db.Prepare(emailExists)
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()
	us, err := u.scanData(email, stmt)
	if err != nil {
		return domain.User{}, err
	}
	return us, nil
}

func (u user) NicknameFieldExists(nickname string) (domain.User, error) {
	stmt, err := u.db.Prepare(nicknameExists)
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()
	us, err := u.scanData(nickname, stmt)
	if err != nil {
		return domain.User{}, err
	}
	return us, nil
}

func (u user) scanData(data interface{},stmt *sql.Stmt) (domain.User, error) {
	us := domain.User{}
	err := stmt.QueryRow(data).Scan(
		us.ID,
		us.Name,
		us.LastName,
		us.Email,
		us.Description,
		us.Picture,
	)
	return us, err
}
