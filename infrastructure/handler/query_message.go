package handler

import (
	"github.com/MikelSot/melody/interfaces"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func (q queryMessage) GetMessageFromChat(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	values := mux.Vars(r)
	myId, err := strconv.Atoi(values["my-id"])
	max, err := strconv.Atoi(values["max"])
	if err != nil || myId < 0 || max < 0{
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}
	messages, err := q.query.GetMessageFromChat(uint(max), uint(myId))
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", messages)
	responseJson(w, http.StatusAccepted, res)
}
