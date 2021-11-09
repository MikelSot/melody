package handler

import (
	"github.com/MikelSot/melody/interfaces"
	"net/http"
)

type persistenceMessage struct {
	pers interfaces.PersistenceMessage
}

func NewPersistenceMessage(pers interfaces.PersistenceMessage) *persistenceMessage {
	return &persistenceMessage{pers}
}

func (p persistenceMessage) Create(w http.ResponseWriter, r *http.Request){

}

func (p persistenceMessage) UpdateStarField(w http.ResponseWriter, r *http.Request){

}

func (p persistenceMessage) DeleteMessage(w http.ResponseWriter, r *http.Request)  {

}