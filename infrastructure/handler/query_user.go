package handler

import (
	"encoding/json"
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/interfaces"
	"net/http"
	"strconv"
	"strings"
)

type queryUser struct {
	query interfaces.QueryUser
}

func NewQueryUser(user interfaces.QueryUser) *queryUser {
	return &queryUser{user}
}

func (q queryUser) GetAllNotAddedUsers(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	max, err := strconv.Atoi(r.URL.Query().Get("max"))
	if err != nil || id < 0 || max < 0{
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	users, err :=q.query.GetAllNotAddedUsers(max, uint(id))
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	res := NewResponse(Message, "ok", users)
	responseJson(w, http.StatusOK, res)
}

func (q queryUser) EmailFieldExists(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	data := domain.CreateUser{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	data.Email = strings.TrimSpace(data.Email)
	if !isEmail(data.Email) {
		res := NewResponse(Error, "Email invalido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	user, err := q.query.EmailFieldExists(data.Email)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	if user.ID != 0{
		res := NewResponse(Error, "Ya existe un usuario con ese email", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	res := NewResponse(Message, "ok", nil)
	responseJson(w, http.StatusOK, res)
}

func (q queryUser) NicknameFieldExists(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	data := domain.FieldString{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	data.Field = strings.TrimSpace(data.Field)
	user, err := q.query.NicknameFieldExists(data.Field)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	if user.ID != 0{
		res := NewResponse(Error, "Ya existe un usuario con ese nickname", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	res := NewResponse(Message, "ok", nil)
	responseJson(w, http.StatusOK, res)
}