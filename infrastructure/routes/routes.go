package routes

import (
	"database/sql"
	"github.com/MikelSot/melody/database/dao/persistence"
	"github.com/MikelSot/melody/infrastructure/middlewares"
	"github.com/MikelSot/melody/interfaces"
	"github.com/gorilla/mux"
)

type router struct {
	mux *mux.Router
	db  *sql.DB
	jwt interfaces.JWT
}

func NewRouter(mux *mux.Router, db *sql.DB, jwt interfaces.JWT) *router {
	return &router{mux, db, jwt}
}

func (r router) User() {
	midd := middlewares.NewMiddleware(r.jwt)
	userPers := newUserRouter(r.mux, midd)
	userdb := persistence.NewUser(r.db)
	userPers.persistenceUser(userdb)
}
