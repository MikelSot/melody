package tables

import "database/sql"

const (
	table_chat_user = `CREATE TABLE IF NOT EXISTS chat_users(
						chat_id BIGINT NOT NULL,
						user_id INT NOT NULL,
						CONSTRAINT pk_chat_users_id PRIMARY KEY (chat_id,user_id),
    					CONSTRAINT fk_chat_users_chat_id FOREIGN KEY (chat_id)
							REFERENCES chats(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    					CONSTRAINT fk_chat_users_user_id FOREIGN KEY (user_id)
							REFERENCES users(id) ON UPDATE RESTRICT ON DELETE RESTRICT
					)`
)

type chatUser struct {
	db *sql.DB
}


func NewChatUsers(db *sql.DB) *chatUser {
	return &chatUser{db}
}

func (c chatUser) Migrate() error {
	stmt, err := c.db.Prepare(table_chat_user)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

