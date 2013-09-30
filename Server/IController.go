package Server

import (
	"net/http"
)

type IController interface {
	CreateController(http.ResponseWriter, *http.Request) IController
}
