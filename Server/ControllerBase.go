package Server

import (
	"net/http"
)

type ControllerBase struct {
	Rw      http.ResponseWriter
	Request *http.Request
}

func (this *ControllerBase) CreateController(rw http.ResponseWriter, request *http.Request) {
	this.Rw = rw
	this.Request = request
}
