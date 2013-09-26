package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Server struct {
	Sites *SiteCollection
}

const (
	SiteConfigFile = "Resource/site.json"
)

var inst *Server

var once sync.Once

func GetCurrentServer() *Server {
	once.Do(newServer)
	return inst
}

func newServer() {
	inst = new(Server)
	inst.init()
}

func (this *Server) init() error {
	line, err := ioutil.ReadFile(SiteConfigFile)
	if err == nil {
		sc := new(SiteCollection)

		tsc := new(tempSiteCollection)

		err := json.Unmarshal(line, tsc)
		if err != nil {
			fmt.Printf(err.Error())
		}

		for i := 0; i < len(tsc.Sites); i++ {
			site := tsc.Sites[i]
			site.init()
			sc.Push(site)
		}

		this.Sites = sc

		return nil
	}
	return err
}

func (this *Server) Start() {
	for _, site := range this.Sites.array {
		go site.Start()
	}
}

func (this *Server) Stop() {
	http.HandleFunc()
}

func (this *Server) defaultHandleFunction(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<h1>Hello World!</br>I'm</h1>")
}

type tempSiteCollection struct {
	Sites []*Site
}
