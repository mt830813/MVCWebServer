package Handler

import (
	"Prj/MVCWebServer/Server"
)

type baseHandler struct {
	site *Server.Site
}

func (this *baseHandler) SetSite(site *Server.Site) {
	this.site = site
}
