package query

import "database/sql"

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *user {
	return &user{db}
}
