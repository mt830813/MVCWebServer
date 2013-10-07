package Controller

import (
	"fmt"
	"github.com/mt830813/MVCWebServer/Server"
)

type TestController struct {
	Server.ControllerBase
}

func (this *TestController) Test(title string, name string) {
	model := new(User)
	model.Name = name
	model.Title = title
	if err := this.View("", model); err != nil {
		fmt.Printf(err.Error())
	}
}

func (this *TestController) Test2(title string, name string) {
	model := new(User)
	model.Name = name
	model.Title = title
	if err := this.View("Test/Test.html", model); err != nil {
		fmt.Printf(err.Error())
	}
}

func (this *TestController) TestObj(user *User) {
	if err := this.View("Test/Test.html", user); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

type User struct {
	Server.ViewModelBase
	Name  string
	Title string
}
