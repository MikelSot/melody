package routes

import (
	"database/sql"
	"github.com/MikelSot/melody/database/dao/persistence"
	"github.com/gorilla/mux"
)

type router struct {
	mux *mux.Router
	db  *sql.DB
}

func NewRouter(mux *mux.Router, db *sql.DB) *router {
	return &router{mux, db}
}

func (r router) User() {
	userPers := newUserRouter(r.mux)
	userdb := persistence.NewUser(r.db)
	userPers.persistenceUser(userdb)
}
