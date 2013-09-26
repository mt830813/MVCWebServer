// MVCWebServer project main.go
package main

import (
	"Prj/MVCWebServer/Common"
	"Prj/MVCWebServer/Server"
	"bufio"
	"fmt"
	//"log"
	"os"
	"reflect"
)

func main() {

	fmt.Printf("ServerStart\n")
	reader := bufio.NewReader(os.Stdin)

	sc := Server.GetCurrentServer()
	TestFactory()
	sc.Start()
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Error:%s", err.Error())
			continue
		}
		command := string(line[:len(line)-2])
		fmt.Printf("Get Command:%s\n", command)
		if command == "stop" {
			break
		}
	}
}

type testType struct {
}

type testInterface interface {
	Test() int
}

func (this *testType) Test() int {
	return 128
}

func TestFactory() {
	factory := Common.GetIOCFactory()

	factory.Regist(reflect.TypeOf((*testInterface)(nil)).Elem(),
		reflect.TypeOf(new(testType)), Common.InstanceType_Normal)

	rResult := new(testType).Test()
	tObj, _ := factory.Get(reflect.TypeOf((*testInterface)(nil)).Elem())

	tResult := tObj.(*testInterface).Test()

	fmt.Printf("%s", rResult == tResult)

}
