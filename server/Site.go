package server

import (
	"fmt"
	"net/http"
)

type Site struct {
	RootPath string
	Port     string
	Name     string
}

func (this *Site) Start() {
	err := http.ListenAndServe(":"+this.Port, this)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}

func (this *Site) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<h1>Hello World!</br>I'm %s</h1>", this.Name)
}
