package Server

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type ControllerBase struct {
	Rw      http.ResponseWriter
	Request *http.Request
	Site    *Site
	Context *RequestContext
}

const (
	viewRelatePath = "/view"
)

func (this *ControllerBase) View(path string, obj interface{}) error {
	model := obj.(IViewModel)
	model = this.initViewModel(model)

	viewPath := this.Site.RootPath + viewRelatePath

	if len([]rune(path)) == 0 {
		viewPath = viewPath + "/" + this.Context.ControllerName + "/" + this.Context.MethodName + ".html"
	} else {
		viewPath = viewPath + "/" + path
	}

	if absViewPath, err := filepath.Abs(viewPath); err != nil {
		return err
	} else {
		if h, e := template.ParseFiles(absViewPath); e == nil {
			err = h.Execute(this.Rw, model)
			return err
		} else {
			return e
		}
	}

}

func (this *ControllerBase) initViewModel(model IViewModel) IViewModel {
	model.SetBottomScript("Bottom")
	model.SetCss("Css")
	model.SetTopScript("Top")
	//todo
	return model
}
