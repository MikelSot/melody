package handler

import (
	"regexp"
)

const (
	regexEmail = `([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9])+\.)+([a-zA-Z0-9]{2,4})`
)

func isEmail(email string) bool {
	r, _ := regexp.Compile(regexEmail)
	if !r.MatchString(email) {
		return false
	}
	return true
}

//TODO: falta las consultas esas seran en http no en websocket, lo unico que se hara con websocke esel chat
