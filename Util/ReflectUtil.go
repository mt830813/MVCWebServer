package Util

import (
	"errors"
	"fmt"
	"reflect"
)

type ReflectUtil struct {
}

func (this *ReflectUtil) RunMethod(obj interface{},
	params []interface{}) ([]interface{}, error) {
	var method reflect.Value
	switch obj.(type) {
	case reflect.Value:
		method = obj.(reflect.Value)
	default:
		if reflect.TypeOf(obj).Kind() == reflect.Func {
			method = reflect.ValueOf(obj)
		} else {
			return nil, errors.New("Only allow run method with func RunMethod")
		}
	}
	argumentsCount := method.Type().NumIn()

	in := make([]reflect.Value, argumentsCount)

	for i := 0; i < argumentsCount; i++ {
		typeIn := method.Type().In(i)
		if i >= len(params) {
			in[i] = reflect.New(typeIn).Elem()
		} else {

			in[i] = reflect.ValueOf(params[i])
		}
	}
	results := method.Call(in)

	returnValue := make([]interface{}, len(results))

	for i, obj := range results {
		returnValue[i] = obj.Interface()
	}

	return returnValue, nil
}

func (this *ReflectUtil) RunDMethod(method interface{},
	args ...interface{}) ([]interface{}, error) {
	return this.RunMethod(method, args)
}

func (this *ReflectUtil) RunDObjMethod(obj interface{}, methodName string,
	params ...interface{}) ([]interface{}, error) {
	return this.RunObjMethod(obj, methodName, params)
}

func (this *ReflectUtil) RunObjMethod(obj interface{},
	methodName string, params []interface{}) ([]interface{}, error) {
	tController := reflect.ValueOf(obj)

	if method := tController.MethodByName(methodName); method.IsValid() {
		return this.RunMethod(method, params)
	} else {
		msg := fmt.Sprintf("run method named<%s> failed", methodName)
		return nil, errors.New(msg)
	}
}
