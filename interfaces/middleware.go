package interfaces

import "net/http"

type Authenticationer interface {
	Authentication(fun http.HandlerFunc) http.HandlerFunc
}
