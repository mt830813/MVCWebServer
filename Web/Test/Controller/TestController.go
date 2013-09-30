package Controller

import (
	"fmt"
	"github.com/mt830813/MVCWebServer/Server"
)

type TestController struct {
	Server.ControllerBase
}

func (this *TestController) Test(title string, name string) {
	if err := this.View("", struct {
		Name  string
		Title string
	}{name, title}); err != nil {
		fmt.Printf(err.Error())
	}
}

func (this *TestController) Test2(title string, name string) {
	if err := this.View("Test/Test.html", struct {
		Name  string
		Title string
	}{name, title}); err != nil {
		fmt.Printf(err.Error())
	}
}
