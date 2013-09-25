package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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
	var uri = request.RequestURI
	var filePath = this.RootPath + uri
	temp, err := ioutil.ReadFile(filePath)
	if err != nil {
		response.WriteHeader(404)
		fmt.Printf("%v~:%s\n", time.Now().Format("2006-01-30 15:04:05"), err.Error())
	} else {
		fmt.Fprint(response, string(temp))
	}
}
