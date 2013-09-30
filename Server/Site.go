package Server

import (
	"Prj/MVCWebServer/Common"

	"fmt"
	"net"
	"net/http"
	"reflect"
	"strings"
	"sync"
	//"time"
)

const (
//strTimeFormat = "2006-01-30 15:04:05"
)

type Site struct {
	RootPath    string
	Port        string
	Name        string
	SiteSetting *SiteSetting

	defaultMux *http.ServeMux
	inst       *http.Server
	l          net.Listener
	once       sync.Once
}

func (this *Site) Start() {
	this.once.Do(this.init)

	l, err := net.Listen("tcp", this.inst.Addr)

	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	this.l = l

	go this.inst.Serve(this.l)

	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	fmt.Printf("site:%s started \n", this.Port)

}

func (this *Site) Stop() {
	/*
		todo
		use this way to stop and restart server always print an AcceptEx to console
		after i did some research, if i want make this acceptex away
		need do many work include implement an interface "net.Listener"
		fuck, why i need cross the fucking wall to visit groups.google.

	*/
	if err := this.l.Close(); err == nil {
		fmt.Printf("site:%s stoped \n", this.Port)
	} else {
		fmt.Printf("site:%s stop error:%s \n", this.Port, err.Error())
	}
}

func (this *Site) init() {
	iControllerType := reflect.TypeOf((*IHandler)(nil)).Elem()
	factory := Common.GetIOCFactory()

	this.defaultMux = http.NewServeMux()
	this.inst = &http.Server{Addr: ":" + this.Port, Handler: this.defaultMux}

	if this.SiteSetting == nil {
		this.SiteSetting = newSiteSetting()
	}

	for _, setting := range this.SiteSetting.Handlers {
		if obj, ok := factory.GetByName(strings.ToLower(setting.Name), iControllerType, nil); ok == nil && obj != nil {
			controller := obj.(IHandler)
			controller.SetSite(this)

			this.defaultMux.Handle(setting.Pattern, controller)
		} else {
			if obj == nil {
				fmt.Printf("handler named<%s> not exist\n", setting.Name)
			} else {
				fmt.Printf("%s\n", ok.Error())
			}
		}
	}
}
