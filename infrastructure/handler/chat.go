package handler

import (
	"encoding/json"
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/interfaces"
	"net/http"
	"strconv"
)

type persistenceChat struct {
	pers     interfaces.PersistenceChater
}

func NewPersistenceChat(pers interfaces.PersistenceChater) *persistenceChat {
	return &persistenceChat{pers}
}

func (p persistenceChat) Create(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	data := domain.Chat{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	p.pers.Create(&data)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", data)
	responseJson(w, http.StatusCreated, res)
}


func (p persistenceChat) Update(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	data := domain.Chat{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	p.pers.Update(&data)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", data)
	responseJson(w, http.StatusOK, res)
}

func (p persistenceChat) GetById(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		res := NewResponse(Error, "id incorrecto", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	chat, err := p.pers.GetById(uint(id))
	if chat.ID == 0 {
		res := NewResponse(Error, "Este chat no existe", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", chat)
	responseJson(w, http.StatusOK, res)
}