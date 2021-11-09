package routes

import (
	"github.com/MikelSot/melody/infrastructure/handler"
	"github.com/MikelSot/melody/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

type chatUser struct {
	mux *mux.Router
	midd interfaces.Authenticationer
}

func newChatUserRouter(mux *mux.Router, midd interfaces.Authenticationer) *chatUser {
	return &chatUser{mux, midd}
}

func (c *chatUser) chatUser(query interfaces.QueryChatUser, pers interfaces.PersistenceChatUser){
	cu := handler.NewChatUser(query,pers)

	c.mux.HandleFunc(
		api+"/chat-user-create",
		c.midd.Authentication(cu.Create),
	).Methods(http.MethodPost)


	c.mux.HandleFunc(
		api+"/chat-user-match/{my-id}/{your-id}",
		c.midd.Authentication(cu.Match),
	).Methods(http.MethodGet)


	c.mux.HandleFunc(
		api+"/chat-user-all-people/{my-id}/{max}",
		c.midd.Authentication(cu.AllPeopleAddedToMyChat),
	).Methods(http.MethodGet)

	c.mux.HandleFunc(
		api+"/chat-user-all-group/{my-id}/{max}",
		c.midd.Authentication(cu.AllGroupsAddedToMyChat),
	).Methods(http.MethodGet)

	c.mux.HandleFunc(
		api+"/chat-user-all/{my-id}/{max}",
		c.midd.Authentication(cu.AllAddedToMyChat),
	).Methods(http.MethodGet)
}