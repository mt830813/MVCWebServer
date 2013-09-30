package Server

import (
	"net/http"
)

type ControllerBase struct {
	Rw      http.ResponseWriter
	Request *http.Request
}
