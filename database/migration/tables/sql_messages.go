package tables

import "database/sql"

const (
	table_message = `CREATE TABLE IF NOT EXISTS messages(
						id BIGSERIAL NOT NULL,
						message VARCHAR(255) DEFAULT '' NOT NULL,
						star BOOLEAN DEFAULT false,
						user_id INT NOT NULL,
						chat_id BIGINT NOT NULL,
						created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
						CONSTRAINT pk_messages_id PRIMARY KEY (id),
						CONSTRAINT fk_message_user_id FOREIGN KEY (user_id)
							REFERENCES users(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
						CONSTRAINT fk_message_chat_id FOREIGN KEY (chat_id)
							REFERENCES chats(id) ON UPDATE RESTRICT ON DELETE RESTRICT
					)`
)

type message struct {
	db *sql.DB
}

func NewMessage(db *sql.DB) *message {
	return &message{db}
}

func (m message) Migrate() error {
	stmt, err := m.db.Prepare(table_message)
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
