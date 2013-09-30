package Handler

import (
	"Prj/MVCWebServer/Common"
	"Prj/MVCWebServer/Server"
	"Prj/MVCWebServer/Util"
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

	tempParmas := args[2:]

	params := make([]interface{}, len(tempParmas))

	for i, arg := range tempParmas {
		params[i] = interface{}(arg)
	}

	context := &Server.RequestContext{ControllerName: controllerName, MethodName: methodName}

	if controller, ok := factory.GetByName(strings.ToLower(controllerName),
		iControllerType, map[string]interface{}{"Rw": w, "Request": r,
			"Site": this.site, "Context": context}); ok != nil || controller == nil {
		w.WriteHeader(404)
		if ok != nil {
			fmt.Printf("view path %s failed:%s\n", requestPath, ok.Error())
		} else {
			fmt.Printf("view path %s failed:controller named <%s> not regist\n", requestPath, controllerName)
		}
	} else {
		if results, err := new(Util.ReflectUtil).RunObjMethod(controller, methodName, params); err != nil {
			fmt.Printf("view path %s failed:%s\n", requestPath, err.Error())
		} else {
			if len(results) > 0 {
				fmt.Fprint(w, results[0])
			}
		}

	}
}
