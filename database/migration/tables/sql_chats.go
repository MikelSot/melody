package tables

import "database/sql"

const (
	table_chat = `CREATE TABLE IF NOT EXISTS chats(
						id BIGSERIAL NOT NULL,
						created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
						CONSTRAINT pk_chats_id PRIMARY KEY (id)
					)`
)

type chat struct {
	db *sql.DB
}

func NewChat(db *sql.DB) *chat {
	return &chat{db}
}

func (m chat) Migrate() error {
	stmt, err := m.db.Prepare(table_chat)
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

