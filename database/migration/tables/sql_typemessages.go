package tables

import "database/sql"

const (
	table_type_message = `CREATE TABLE IF NOT EXISTS type_messages(
							id SMALLSERIAL NOT NULL,	
							type VARCHAR(30) NOT NULL,
							CONSTRAINT pk_type_messages_id PRIMARY KEY (id),
							CONSTRAINT un_type_messages_type UNIQUE (type)
						)`
)

type typeMessage struct {
	db *sql.DB
}

func NewTypeMessage(db *sql.DB) *typeMessage {
	return &typeMessage{db}
}

func (m typeMessage) Migrate() error {
	stmt, err := m.db.Prepare(table_type_message)
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