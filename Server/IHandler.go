package Server

import (
	"net/http"
)

type IHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)

	SetSite(*Site)
}
