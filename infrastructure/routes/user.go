package routes

import (
	"github.com/MikelSot/melody/infrastructure/handler"
	"github.com/MikelSot/melody/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	api = "/api/v1"
)

type user struct {
	mux *mux.Router
	midd interfaces.Authenticationer
}

func newUserRouter(mux *mux.Router, midd interfaces.Authenticationer) *user{
	return &user{mux, midd}
}

func (u *user) loginAndGetById(query interfaces.QueryUser, jwt interfaces.JWT){
	login := handler.NewLogin(query,jwt)
	u.mux.HandleFunc(api+"/login", login.Login).Methods(http.MethodPost)
	u.mux.HandleFunc(api+"/get-id/{id}",u.midd.Authentication(login.GetUserById)).Methods(http.MethodGet)
}


func (u user) persistenceUser(user interfaces.PersistenceUser){
	pers := handler.NewPersistenceUser(user)

	u.mux.HandleFunc(
		api+"/user-register",
		pers.Create,
	).Methods(http.MethodPost)

	u.mux.HandleFunc(
		api+"/user-update/{id}",
		u.midd.Authentication(pers.Update),
	).Methods(http.MethodPut)

	u.mux.HandleFunc(
		api+"/user-update-nickname/{id}",
		u.midd.Authentication(pers.UpdateNicknameField),
	).Methods(http.MethodPut)

	u.mux.HandleFunc(
		api+"/user-update-picture/{id}",
		u.midd.Authentication(pers.UpdatePictureField),
	).Methods(http.MethodPut)

	u.mux.HandleFunc(
		api+"/user-update-online/{id}",
		u.midd.Authentication(pers.UpdateOnlineField),
	).Methods(http.MethodPut)
}

func (u user) queryUser(query interfaces.QueryUser)  {
	 que := handler.NewQueryUser(query)

	u.mux.HandleFunc(
		api+"/user-update-online/{id}/{max}",
		u.midd.Authentication(que.GetAllNotAddedUsers),
	).Methods(http.MethodGet)

	u.mux.HandleFunc(
		api+"/user-email-exists",
		u.midd.Authentication(que.EmailFieldExists),
	).Methods(http.MethodPost)

	u.mux.HandleFunc(
		api+"/user-nickname-exists",
		u.midd.Authentication(que.NicknameFieldExists),
	).Methods(http.MethodPost)
}
