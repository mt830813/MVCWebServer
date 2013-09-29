package Command

import (
	"Prj/MVCWebServer/Server"
)

type Stop struct {
	*CommandBase
}

func (this *Stop) DoCommand(param string) {
	server := Server.GetCurrentServer()
	server.Stop()
}
