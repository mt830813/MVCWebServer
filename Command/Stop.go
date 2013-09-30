package Command

import (
	"github.com/mt830813/MVCWebServer/Server"
)

type Stop struct {
	*CommandBase
}

func (this *Stop) DoCommand(param string) {
	server := Server.GetCurrentServer()
	server.Stop()
}
