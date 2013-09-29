package Command

import (
	"fmt"
)

type Echo struct {
	*CommandBase
}

func (this *Echo) DoCommand(param string) {
	fmt.Println("Echo:" + param)
}
