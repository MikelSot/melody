package handler

import (
	"encoding/json"
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/interfaces"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type persistenceUser struct {
	pers interfaces.PersistenceUser
}

func NewPersistenceUser(pers interfaces.PersistenceUser) *persistenceUser {
	return &persistenceUser{pers}
}

func (p persistenceUser) Create(w http.ResponseWriter, r *http.Request)  {
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

	data.Password = p.encrypt(data.Password)
	err = p.pers.Create(&data)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}

	id := map[string]uint{"id":data.ID}
	res := NewResponse(Message, "ok", id)
	responseJson(w, http.StatusCreated, res)
}

func (p persistenceUser) Update(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPut {
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
	data := domain.UpdateUser{}; data.ID= uint(id)
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}
	err = p.pers.Update(&data)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}
	res := NewResponse(Message, "ok", nil)
	responseJson(w, http.StatusOK, res)
}

func (p persistenceUser) UpdateNicknameField(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPut {
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
	data := domain.FieldString{}; data.ID= uint(id)
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	err =p.pers.UpdateNicknameField(&data)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}
	res := NewResponse(Message, "ok", nil)
	responseJson(w, http.StatusOK, res)
}

func (p persistenceUser) UpdatePictureField(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPut {
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
	data := domain.FieldString{}; data.ID= uint(id)
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	err =p.pers.UpdatePictureField(&data)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}
	res := NewResponse(Message, "ok", nil)
	responseJson(w, http.StatusOK, res)
}

func (p persistenceUser) UpdateOnlineField(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut {
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
	data := domain.State{}; data.ID= uint(id)
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	err =p.pers.UpdateOnlineField(&data)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}
	res := NewResponse(Message, "ok", nil)
	responseJson(w, http.StatusOK, res)
}

func (p persistenceUser) encrypt(pass string) string {
	cost := 6 // es el numero de veces que recorre y encripta
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		log.Println("Error al encriptar contraseña")
	}
	return string(bytes)
}