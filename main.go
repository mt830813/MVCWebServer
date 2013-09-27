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

type testTypeDecorater struct {
	innerPackage testInterface
}

func (this *testTypeDecorater) Test() int {
	return this.innerPackage.Test() * 4
}

func (this *testTypeDecorater) SetPackage(i interface{}) {
	obj := i.(testInterface)
	this.innerPackage = obj
}

func TestFactory() {
	factory := Common.GetIOCFactory()

	ti := reflect.TypeOf((*testInterface)(nil)).Elem()

	factory.Regist(reflect.TypeOf((*testInterface)(nil)).Elem(),
		reflect.TypeOf(new(testType)), Common.InstanceType_Normal)
	factory.RegistDecorate(ti, reflect.TypeOf(new(testTypeDecorater)), Common.InstanceType_Singleton)

	rResult := new(testType).Test()
	tObj, _ := factory.Get(ti)

	tResult := tObj.(testInterface).Test()

	fmt.Printf("%s", rResult == tResult)

}
