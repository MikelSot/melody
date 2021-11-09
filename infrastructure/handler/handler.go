package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
)

const (
	Error      = "error"
	Message    = "message"
	regexEmail = `([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9])+\.)+([a-zA-Z0-9]{2,4})`
)


type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func NewResponse(messageType string, message string, data interface{}) response {
	return response{
		messageType,
		message,
		data,
	}
}

func responseJson(w http.ResponseWriter, statusCode int,r response){
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isEmail(email string) bool {
	r, _ := regexp.Compile(regexEmail)
	if !r.MatchString(email) {
		return false
	}
	return true
}