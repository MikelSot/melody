package routes

import (
	"github.com/MikelSot/melody/infrastructure/handler"
	"github.com/MikelSot/melody/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

type chat struct {
	mux *mux.Router
	midd interfaces.Authenticationer
}

func newChatRouter(mux *mux.Router, midd interfaces.Authenticationer) *chat{
	return &chat{mux, midd}
}

func (c chat) persistenceChat(chat interfaces.PersistenceChater) {
	pers := handler.NewPersistenceChat(chat)

	c.mux.HandleFunc(
		api+"/chat-create",
		c.midd.Authentication(pers.Create),
	).Methods(http.MethodPost)

	c.mux.HandleFunc(
		api+"/chat-get-by-id/{id}",
		c.midd.Authentication(pers.GetById),
	).Methods(http.MethodGet)
}