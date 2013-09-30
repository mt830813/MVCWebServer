package Controller

import (
	"Prj/MVCWebServer/Server"
	"fmt"
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
