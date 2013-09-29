package Command

import (
	"Prj/MVCWebServer/Common"
	"fmt"
	"reflect"
)

type List struct {
	*CommandBase
}

func (this *List) DoCommand(param string) {
	fmt.Printf("System contains blow Command\n")

	factory := Common.GetIOCFactory()
	iBaseType := reflect.TypeOf((*ICommand)(nil)).Elem()

	array := factory.GetAll(iBaseType)

	index := 0

	for key, _ := range array {
		if index > 0 && index%4 == 0 {
			fmt.Println()
		}
		index++
		fmt.Printf("%s\t", key)
	}
	fmt.Println()
}

func (this *List) GetName() string {
	return "List"
}

func (this *List) GetHelp() string {
	return "to list all the command registed in this system"
}
