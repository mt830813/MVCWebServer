package Command

import (
	"Prj/MVCWebServer/Server"
)

type Start struct {
	*CommandBase
}

func (this *Start) DoCommand(param string) {
	server := Server.GetCurrentServer()
	server.Start()
}
