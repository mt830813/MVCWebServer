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
	inst.array = make(map[string]interfaceArrayValue)
}

func (this *IOCFactory) Regist(i reflect.Type, t reflect.Type,
	instType InstanceType) error {
	return this.RegistByName("default", i, t, instType)
}

func (this *IOCFactory) RegistByName(key string, i reflect.Type,
	t reflect.Type, instType InstanceType) error {
	if !this.checkIsImplementInterface(i, t) {
		return fmt.Errorf("regist type error")
	}
	var pArray = this.getPArray(i)
	pArray[key] = this.createNormalRegistContext(i, t, instType)
	return nil
}

func (this *IOCFactory) RegistDecorate(i reflect.Type, t reflect.Type,
	instType InstanceType) {
	this.RegistDecorateByName("default", i, t, instType)
}

func (this *IOCFactory) RegistDecorateByName(key string, i reflect.Type,
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
		idType := reflect.TypeOf((*IDecorater)(nil)).Elem()
		if !this.checkIsImplementInterface(idType, t) {
			log.Printf("strcut not implement interface IDecorater can't regist as a decorater")
			return
		}
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

func (this *IOCFactory) Get(i reflect.Type) (interface{}, error) {
	return this.GetByName("default", i)
}

func (this *IOCFactory) GetByName(key string, i reflect.Type) (interface{}, error) {
	var returnValue interface{}
	if iContext, err := this.getRegistContext(key, i); err != nil {
		return nil, err
	} else {
		switch iContext.(type) {
		case *registContext:
			regContext := iContext.(*registContext)
			returnValue = this.createNewInst(regContext)
			if regContext.instType == InstanceType_Singleton {
				pArray := this.getPArray(i)
				pArray[key] = returnValue
			}
		case *decorateRegistcontext:
			drContext := iContext.(*decorateRegistcontext)
			returnValue = this.createNewDecorateInst(drContext)
		default:
			returnValue = iContext
		}
	}
	return returnValue, nil

}

func (this *IOCFactory) getRegistContext(key string, i reflect.Type) (interface{}, error) {
	var pArray = this.getPArray(i)
	if len(pArray) == 0 {
		return nil, fmt.Errorf("interface named \"%s\" not regist any type", reflect.TypeOf(i).Name())
	}

	return pArray[key], nil
}

func (this *IOCFactory) getPArray(i reflect.Type) interfaceArrayValue {

	pName := i.Name()
	if this.array[pName] == nil {
		this.array[pName] = make(interfaceArrayValue)
	}
	return this.array[pName]
}

func (this *IOCFactory) createNormalRegistContext(i reflect.Type,
	t reflect.Type, instType InstanceType) *registContext {
	returnValue := new(registContext)
	returnValue.bType = t
	returnValue.instType = instType
	return returnValue
}

func (this *IOCFactory) checkIsImplementInterface(i reflect.Type, instType reflect.Type) bool {
	return instType.Implements(i)
}

func (this *IOCFactory) createNewInst(context *registContext) interface{} {
	return reflect.New(context.bType).Elem().Interface()
}

func (this *IOCFactory) createNewDecorateInst(context *decorateRegistcontext) interface{} {
	returnValue := this.createNewInst(context.currentContext)
	fmt.Printf("context:%v\n", context.nextContext)
	if context.nextContext != nil {
		returnValue.(IDecorater).SetPackage(this.createNewDecorateInst(context.nextContext))
	}
	return returnValue
}
