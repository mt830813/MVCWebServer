package Command

import (
	"github.com/mt830813/MVCWebServer/Server"
)

type Start struct {
	*CommandBase
}

func (this *Start) DoCommand(param string) {
	server := Server.GetCurrentServer()
	server.Start()
}
