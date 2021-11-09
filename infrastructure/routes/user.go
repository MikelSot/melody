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
}

func newUserRouter(mux *mux.Router) *user{
	return &user{mux}
}

func (u *user) loginAndGetById(query interfaces.QueryUser, jwt interfaces.JWT){
	login := handler.NewLogin(query,jwt)
	u.mux.HandleFunc(api+"/login", login.Login).Methods(http.MethodPost)
	u.mux.HandleFunc(api+"/get-id/{id}", login.GetUserById).Methods(http.MethodGet)
}


func (u *user) persistenceUser(user interfaces.PersistenceUser){
	pers := handler.NewPersistenceUser(user)
	u.mux.HandleFunc(api+"/register", pers.Create).Methods(http.MethodPost)
	u.mux.HandleFunc(api+"/update/{id}", pers.Update).Methods(http.MethodPut)
	u.mux.HandleFunc(api+"/update-nickname/{id}", pers.UpdateNicknameField).Methods(http.MethodPut)
	u.mux.HandleFunc(api+"/update-picture/{id}", pers.UpdatePictureField).Methods(http.MethodPut)
	u.mux.HandleFunc(api+"/update-online/{id}", pers.UpdateOnlineField).Methods(http.MethodPut)
}

