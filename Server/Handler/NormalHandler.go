package Handler

import (
	"Prj/MVCWebServer/Common"
	"Prj/MVCWebServer/Server"
	"fmt"
	"net/http"

	"reflect"
	"strings"
)

type NormalHandler struct {
	baseHandler
}

func (this *NormalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iControllerType := reflect.TypeOf((*Server.IController)(nil)).Elem()

	requestPath := strings.TrimLeft(r.URL.Path, "/")
	args := strings.Split(requestPath, "/")

	factory := Common.GetIOCFactory()

	if len(args) < 2 {
		w.WriteHeader(404)
		return
	}

	controllerName := args[0]

	methodName := args[1]
	params := args[2:]

	if controller, ok := factory.GetByName(strings.ToLower(controllerName), iControllerType); ok != nil || controller == nil {
		w.WriteHeader(404)
		if ok != nil {
			fmt.Printf("view path %s failed:%s\n", requestPath, ok.Error())
		} else {
			fmt.Printf("view path %s failed:controller named <%s> not regist\n", requestPath, controllerName)
		}
	} else {
		tController := reflect.ValueOf(controller)

		if method := tController.MethodByName(methodName); method.IsValid() {

			argumentsCount := method.Type().NumIn()

			in := make([]reflect.Value, argumentsCount)

			for i := 0; i < argumentsCount; i++ {
				if i >= len(params) {
					typeIn := method.Type().In(i)
					in[i] = reflect.New(typeIn).Elem()
				} else {
					in[i] = reflect.ValueOf(params[i])
				}
			}
			results := method.Call(in)

			fmt.Fprintf(w, "%s", results[0])
		} else {
			w.WriteHeader(404)
			fmt.Printf("view path %s failed: not contain method named <%s> in controller <%s>\n",
				requestPath, methodName, controllerName)
		}
	}
}
