package Util

import (
	"errors"
	"fmt"
	"reflect"
)

type ReflectUtil struct {
}

func (this *ReflectUtil) RunObjMapMethod(obj interface{}, methodName string, params map[string]interface{}) ([]interface{}, error) {
	tController := reflect.ValueOf(obj)

	if method := tController.MethodByName(methodName); method.IsValid() {
		return this.RunMapMethod(method, params)
	} else {
		msg := fmt.Sprintf("run method named<%s> failed", methodName)
		return nil, errors.New(msg)
	}
}

func (this *ReflectUtil) RunMapMethod(function interface{}, params map[string]interface{}) ([]interface{}, error) {
	method, err := this.getMethod(function)
	if err != nil {
		return nil, err
	}
	mType := method.Type()
	count := mType.NumIn()
	if count != 1 {
		return nil, errors.New("parameter count errer\n")
	}

	if mType.In(0).Kind() != reflect.Ptr {
		return nil, errors.New("parameter kind errer\n")
	}

	newObj := this.typeSetFields(mType.In(0), params)

	return this.RunDMethod(function, newObj)
}

func (this *ReflectUtil) RunMethod(obj interface{},
	params []interface{}) ([]interface{}, error) {
	method, err := this.getMethod(obj)

	if err != nil {
		return nil, err
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

func (this *ReflectUtil) ObjSetFields(objType reflect.Value, values map[string]interface{}) {
	count := objType.NumField()

	for key, value := range values {
		field := objType.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			fType := field.Type()

			if fType.Kind() == reflect.Ptr {
				cParams := value.(map[string]interface{})
				cObj := this.typeSetFields(fType, cParams)
				field.Set(reflect.ValueOf(cObj))
			} else {
				field.Set(reflect.ValueOf(value))
			}
		}
	}
}

func (this *ReflectUtil) typeSetFields(objType reflect.Type, values map[string]interface{}) interface{} {
	newObj := reflect.New(objType.Elem())

	this.ObjSetFields(newObj.Elem(), values)

	return newObj.Interface()
}

func (this *ReflectUtil) getMethod(obj interface{}) (reflect.Value, error) {
	var method reflect.Value
	switch obj.(type) {
	case reflect.Value:
		method = obj.(reflect.Value)
	default:
		if reflect.TypeOf(obj).Kind() == reflect.Func {
			method = reflect.ValueOf(obj)
		} else {
			return reflect.ValueOf(nil), errors.New("Only allow run method with func RunMethod")
		}
	}
	return method, nil
}
