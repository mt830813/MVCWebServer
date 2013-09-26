package Common

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

type InstanceType int

var once sync.Once

const (
	InstanceType_Singleton InstanceType = 1 << iota
	InstanceType_Normal
)

type interfaceArrayValue map[string]typeArrayValue

type typeArrayValue interface{}

type IOCFactory struct {
	array map[string]interfaceArrayValue
}

var inst *IOCFactory

func GetIOCFactory() *IOCFactory {
	once.Do(initFactory)
	return inst
}

func initFactory() {
	inst = new(IOCFactory)

}

func (this *IOCFactory) Regist(i interface{}, t reflect.Type,
	instType InstanceType) {
	this.RegistByName("default", i, t, instType)
}

func (this *IOCFactory) RegistByName(key string, i interface{},
	t reflect.Type, instType InstanceType) {

	var pArray = this.getPArray(i)
	pArray[key] = this.createNormalRegistContext(i, t, instType)
}

func (this *IOCFactory) RegistDecorate(i interface{}, t reflect.Type,
	instType InstanceType) {
	this.RegistDecorateByName("default", i, t, instType)
}

func (this *IOCFactory) RegistDecorateByName(key string, i interface{},
	t reflect.Type, instType InstanceType) {
	pArray := this.getPArray(i)
	rContext, err := this.getRegistContext(key, i)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	dContext := new(decorateRegistcontext)
	dContext.currentContext = this.createNormalRegistContext(i, t, instType)
	var cContext *decorateRegistcontext
	if rContext != nil {
		switch rContext.(type) {
		case *registContext:
			cContext = new(decorateRegistcontext)
			cContext.currentContext = rContext.(*registContext)
		case *decorateRegistcontext:
			cContext = rContext.(*decorateRegistcontext)
		}
		dContext.nextContext = cContext
	}
	pArray[key] = dContext
}

func (this *IOCFactory) getRegistContext(key string, i interface{}) (interface{}, error) {
	var pArray = this.getPArray(i)
	if len(pArray) == 0 {
		return nil, fmt.Errorf("interface named \"%s\" not regist any type", reflect.TypeOf(i).Name())
	}

	return pArray[key], nil
}

func (this *IOCFactory) getPArray(i interface{}) interfaceArrayValue {
	pType := reflect.TypeOf(i)
	pName := pType.Name()
	if this.array[pName] == nil {
		this.array[pName] = make(interfaceArrayValue, 0)
	}
	return this.array[pName]
}

func (this *IOCFactory) createNormalRegistContext(i interface{},
	t reflect.Type, instType InstanceType) *registContext {
	returnValue := new(registContext)
	returnValue.bType = t
	returnValue.instType = instType
	return returnValue
}

func (this *IOCFactory) createNewInst(context *registContext) interface{} {
	//to do
	return nil
}
