package handler

import (
	"github.com/MikelSot/melody/interfaces"
	"net/http"
)

type queryMessage struct {
	query     interfaces.QueryMessage
}

func NewMessage(query interfaces.QueryMessage) *queryMessage {
	return &queryMessage{query}
}

func (q queryMessage)GetAllStarred(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	messages, err := q.query.GetAllStarred()
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}
	res := NewResponse(Message, "ok",messages)
	responseJson(w, http.StatusOK, res)
}

