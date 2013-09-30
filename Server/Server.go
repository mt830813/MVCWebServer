package Server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"sync"
)

type Server struct {
	Sites     *SiteCollection
	isRunning bool
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

	err := inst.init()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (this *Server) Reload() {
	if this.isRunning {
		fmt.Printf("server stopping ...\n")
		this.Stop()
		fmt.Printf("server stopping finish ...\n")
	}
	fmt.Printf("server reload ...\n")
	this.init()
	fmt.Printf("server reload finish ...\n")

	this.Start()
	fmt.Printf("server started ...\n")
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
			sc.Push(site)
		}

		this.Sites = sc

		return nil
	}
	return err
}

func (this *Server) Start() {
	for _, site := range this.Sites.array {
		site.Start()
	}
	this.isRunning = true
}

func (this *Server) Stop() {
	for _, site := range this.Sites.array {
		site.Stop()
	}
	this.isRunning = false
}

func (this *Server) defaultHandleFunction(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<h1>Hello World!</br>I'm</h1>")
}

type tempSiteCollection struct {
	Sites []*Site
}
