// MVCWebServer project main.go
package main

import (
	"Prj/MVCWebServer/Command"
	"Prj/MVCWebServer/Common"
	"Prj/MVCWebServer/Server"
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const ()

func main() {
	factory := Common.GetIOCFactory()

	iCmd := reflect.TypeOf((*ICommand)(nil)).Elem()

	fmt.Printf("ServerStart\n")
	reader := bufio.NewReader(os.Stdin)

	sc := Server.GetCurrentServer()
	sc.Start()

	registCommand()

	for {
		fmt.Print("cmd:")
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Error:%s", err.Error())
			continue
		}
		args := string(line[:len(line)-2])

		var strCommand string

		var strParam string

		if spaceIndex := strings.Index(args, " "); spaceIndex >= 0 {
			runeArgs := []rune(args)
			strCommand = string(runeArgs[:spaceIndex])
			strParam = string(runeArgs[spaceIndex:])
		} else {
			strCommand = args
		}

		strCommand = strings.ToLower(strCommand)

		strParam = strings.Trim(strParam, " ")

		if obj, err := factory.GetByName(strCommand, iCmd); err != nil || obj == nil {
			fmt.Printf("command<%s> not exist in system\n", strCommand)
			continue
		} else {
			command := obj.(ICommand)
			command.DoCommand(strParam)
		}
	}
}

func registCommand() {
	iCmd := reflect.TypeOf((*ICommand)(nil)).Elem()

	factory := Common.GetIOCFactory()

	factory.RegistByName("stop", iCmd, reflect.TypeOf(new(Command.Stop)), Common.InstanceType_Singleton)
	factory.RegistByName("echo", iCmd, reflect.TypeOf(new(Command.Echo)), Common.InstanceType_Singleton)
}
