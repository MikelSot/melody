package middlewares

import (
	"github.com/MikelSot/melody/interfaces"
	"net/http"
)

type middleware struct {
	jwt interfaces.JWT
}

func NewMiddleware(jwt interfaces.JWT) *middleware {
	return &middleware{jwt}
}

func (m middleware) Authentication(fun http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := m.jwt.ValidateToken(token)
		if err != nil {
			m.forbidden(w,r)
			return
		}
		fun(w, r)
	}
}

func (m middleware) forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorizacion"))
}