package main

import (
	"Prj/MVCWebServer/Command"
	"Prj/MVCWebServer/Common"
	"Prj/MVCWebServer/Server"
	"Prj/MVCWebServer/Server/Handler"
	"Prj/MVCWebServer/Web/Test/Controller"
	"reflect"
)

type typeRegist struct {
}

func (this *typeRegist) Regist() {
	registCommand()
	registHandler()
	registController()
}

func registHandler() {
	factory := Common.GetIOCFactory()

	iHandlerType := reflect.TypeOf((*Server.IHandler)(nil)).Elem()

	factory.RegistByName("normalhandler", iHandlerType, reflect.TypeOf(new(Handler.NormalHandler)), Common.InstanceType_Normal, nil)
}

func registCommand() {
	iCmd := reflect.TypeOf((*Command.ICommand)(nil)).Elem()

	factory := Common.GetIOCFactory()

	factory.RegistByName("stop", iCmd, reflect.TypeOf(new(Command.Stop)), Common.InstanceType_Singleton, nil)
	factory.RegistByName("echo", iCmd, reflect.TypeOf(new(Command.Echo)), Common.InstanceType_Singleton, nil)
	factory.RegistByName("start", iCmd, reflect.TypeOf(new(Command.Start)), Common.InstanceType_Singleton, nil)
	factory.RegistByName("list", iCmd, reflect.TypeOf(new(Command.List)), Common.InstanceType_Singleton, nil)
	factory.RegistByName("exit", iCmd, reflect.TypeOf(new(Command.Exit)), Common.InstanceType_Singleton, nil)
	factory.RegistByName("quit", iCmd, reflect.TypeOf(new(Command.Exit)), Common.InstanceType_Singleton, nil)
}

func registController() {
	i := reflect.TypeOf((*Server.IController)(nil)).Elem()

	factory := Common.GetIOCFactory()

	factory.RegistByName("test", i, reflect.TypeOf(new(Controller.TestController)), Common.InstanceType_Normal, nil)

}
