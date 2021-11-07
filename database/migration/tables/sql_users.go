package tables

import "database/sql"

const (
	table_user = `CREATE TABLE IF NOT EXISTS users(
					id SERIAL NOT NULL,
					name VARCHAR(80) DEFAULT '' NOT NULL,
					last_name VARCHAR(80) DEFAULT '' NOT NULL,
					nick_name VARCHAR(80) DEFAULT '' NOT NULL,
					email VARCHAR(150) NOT NULL,
					password VARCHAR(200) NOT NULL,
					online BOOLEAN DEFAULT false,
					description VARCHAR(150) DEFAULT '' NOT NULL,
					picture VARCHAR(200),
					created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
					deleted_at TIMESTAMP WITH TIME ZONE,
					CONSTRAINT pk_users_id PRIMARY KEY (id),
					CONSTRAINT un_users_email UNIQUE (email)
				)`
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *user {
	return &user{db}
}

func (u user) Migrate() error {
	stmt, err := u.db.Prepare(table_user)
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
