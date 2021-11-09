package handler

import (
	"encoding/json"
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/interfaces"
	"net/http"
	"strconv"
)

type chatUser struct {
	query interfaces.QueryChatUser
	pers interfaces.PersistenceChatUser
}

func NewChatUser(user interfaces.QueryChatUser, pers interfaces.PersistenceChatUser) *chatUser {
	return &chatUser{user, pers}
}

func (c chatUser) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	data := domain.ChatUser{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	if data.ChatID <= 0 || data.UserID <= 0 {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	err = c.pers.Create(&data)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", data)
	responseJson(w, http.StatusCreated, res)
}

func (c chatUser) Match(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	myId, err := strconv.Atoi(r.URL.Query().Get("my-id"))
	yourId, err := strconv.Atoi(r.URL.Query().Get("your-id"))
	if err != nil || myId < 0 || yourId < 0{
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}


	match, err := c.query.Match(uint(myId), uint(yourId))
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	if match.ChatID != 0 && match.UserID != 0 {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	res := NewResponse(Message, "ok", match)
	responseJson(w, http.StatusAccepted, res)
}

func (c chatUser) AllPeopleAddedToMyChat(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	myId, err := strconv.Atoi(r.URL.Query().Get("my-id"))
	max, err := strconv.Atoi(r.URL.Query().Get("max"))
	if err != nil || myId < 0 || max < 0{
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	chatusers, err := c.query.AllPeopleAddedToMyChat(uint(max), uint(myId))
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", chatusers)
	responseJson(w, http.StatusAccepted, res)
}

func (c chatUser) AllGroupsAddedToMyChat(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	myId, err := strconv.Atoi(r.URL.Query().Get("my-id"))
	max, err := strconv.Atoi(r.URL.Query().Get("max"))
	if err != nil || myId < 0 || max < 0{
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}


	chatusers, err := c.query.AllPeopleAddedToMyChat(uint(max), uint(myId))
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", chatusers)
	responseJson(w, http.StatusAccepted, res)
}

func (c chatUser) AllAddedToMyChat(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	myId, err := strconv.Atoi(r.URL.Query().Get("my-id"))
	max, err := strconv.Atoi(r.URL.Query().Get("max"))
	if err != nil || myId < 0 || max < 0{
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}


	chatusers, err := c.query.AllPeopleAddedToMyChat(uint(max), uint(myId))
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", chatusers)
	responseJson(w, http.StatusAccepted, res)
}
