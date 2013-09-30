package Controller

import (
	"Prj/MVCWebServer/Server"
	"fmt"
)

type TestController struct {
	Server.ControllerBase
}

func (this *TestController) Test(title string, name string) string {
	returnValue := fmt.Sprintf("<html><body><h1>hi:%s</h1>title:%s</br>name:%s</body></html>", this.Request.URL.Path,
		title, name)
	return returnValue
}
