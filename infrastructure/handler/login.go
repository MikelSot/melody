package handler

import (
	"encoding/json"
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/interfaces"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
)

type login struct {
	exists interfaces.QueryUser
	jwt     interfaces.JWT
}

func NewLogin(e interfaces.QueryUser, j interfaces.JWT) *login {
	return &login{e, j}
}

func (l *login) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := NewResponse(Error, "Metodo no permitido", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	data := domain.Login{}
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

	user, err := l.exists.EmailFieldExists(data.Email)
	if err != nil {
		res := NewResponse(Error, "Email incorrecto", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	pass := []byte(data.Password)
	passDB := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(passDB, pass)
	if err != nil {
		res := NewResponse(Error, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	token, err := l.jwt.GenerateToken(user.ID)
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}
	dataToken := map[string]string{"token":token}
	res := NewResponse(Message, "ok", dataToken)
	responseJson(w, http.StatusOK, res)
}

func (l *login) GetUserById(w http.ResponseWriter, r *http.Request) {
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

	user ,err := l.exists.GetById(uint(id))
	if err != nil {
		res := NewResponse(Error, "Ocurrió un error", nil)
		responseJson(w, http.StatusInternalServerError, res)
		return
	}
	user.Password = ""
	res := NewResponse(Message, "ok", user)
	responseJson(w, http.StatusOK, res)
}
