package routes

import (
	"database/sql"
	"github.com/MikelSot/melody/database/dao/persistence"
	"github.com/MikelSot/melody/database/dao/query"
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
	userque := query.NewUser(r.db)
	userdb := persistence.NewUser(r.db)
	userRouter := newUserRouter(r.mux, midd)

	userRouter.loginAndGetById(userque, r.jwt)
	userRouter.persistenceUser(userdb)
	userRouter.queryUser(userque)
}

func (r router) Chat()  {

}

func (r router) ChatUser() {

}

func (r router) Message()  {

}

