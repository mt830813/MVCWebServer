package Command

import (
	"Prj/MVCWebServer/Server"
	"os"
)

type Stop struct {
}

func (this *Stop) DoCommand(param string) {
	server := Server.GetCurrentServer()
	server.Stop()
	os.Exit(0)
}
