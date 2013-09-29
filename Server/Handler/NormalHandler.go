package Handler

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type NormalHandler struct {
	baseHandler
}

func (this *NormalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rootPath := this.site.RootPath
	controllerPath := rootPath + "/Controller"

	absPath, _ := filepath.Abs(controllerPath)

	fmt.Fprintf(w, "<html><body><h1>%s</h1></body></html>", absPath)
}
