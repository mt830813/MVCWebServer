package Controller

import (
	"Prj/MVCWebServer/Server"
)

type TestController struct {
	Server.ControllerBase
}

func (this *TestController) Test(title string, name string) string {
	return "<html><body><h1>hi</h1>title:" + title + "</br>name:" + name + "</body></html>"
}
