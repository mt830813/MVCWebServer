package Handler

import (
	"github.com/mt830813/MVCWebServer/Server"
)

type baseHandler struct {
	site *Server.Site
}

func (this *baseHandler) SetSite(site *Server.Site) {
	this.site = site
}
