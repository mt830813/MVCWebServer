package Command

import (
	"os"
)

type Exit struct {
	*CommandBase
}

func (this *Exit) DoCommand(param string) {
	os.Exit(0)
}
