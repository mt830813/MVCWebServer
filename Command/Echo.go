package Command

import (
	"fmt"
)

type Echo struct {
}

func (this *Echo) DoCommand(param string) {
	fmt.Println("Echo:" + param)
}
