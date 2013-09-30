package Command

import (
	"github.com/mt830813/MVCWebServer/Server"
)

type ServerReload struct {
	*CommandBase
}

func (this *ServerReload) DoCommand(param string) {
	server := Server.GetCurrentServer()
	server.Reload()
}
