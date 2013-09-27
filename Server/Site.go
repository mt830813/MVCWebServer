package Server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//	"reflect"
	"time"
)

const (
	strTimeFormat = "2006-01-30 15:04:05"
)

type Site struct {
	RootPath      string
	Port          string
	Name          string
	Handles       []string
	_handleModels []IWebHandler
	inst          *http.Server
}

func (this *Site) Start() {
	this.inst = &http.Server{Addr: ":" + this.Port, Handler: this}

	err := this.inst.ListenAndServe()

	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}

func (this *Site) Stop() {

}

func (this *Site) init() {
}

func (this *Site) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	var uri = request.RequestURI
	var filePath = this.RootPath + uri

	temp, err := ioutil.ReadFile(filePath)

	if this._handleModels != nil {
		for _, handler := range this._handleModels {
			result := handler.Handle(response, request)

			if result.isShowMsg {
				fmt.Printf("%s~:%s", time.Now().Format(strTimeFormat), result.msg)
			}

			if result.isSuspand {
				return
			}
		}
	}
	if err != nil {
		response.WriteHeader(404)
		fmt.Printf("%v~:%s\n", time.Now().Format(strTimeFormat), err.Error())
	} else {
		fmt.Fprint(response, string(temp))
	}
}
