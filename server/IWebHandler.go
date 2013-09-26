package Server

import (
	"net/http"
)

type IWebHandler interface {
	Handle(http.ResponseWriter, *http.Request) HandleResult
}
