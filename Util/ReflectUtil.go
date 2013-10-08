package Util

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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

// objType must be the element's value not the pointer's value
func (this *ReflectUtil) ObjSetFields(objType reflect.Value, values map[string]interface{}) {
	for key, value := range values {
		this.objSetFieldByName(objType, key, value)
	}
}

func (this *ReflectUtil) objSetFieldByName(objValue reflect.Value, fieldName string, value interface{}) {
	getArray := func(str string) (is bool, name string, index int) {
		strExp := "(.*)\\[(\\d+)\\]"
		exp := regexp.MustCompile(strExp)
		result := exp.FindStringSubmatch(str)
		is = result != nil
		if is {
			name = result[1]
			index, _ = strconv.Atoi(result[2])
		}
		return
	}

	setValue := func(field reflect.Value, value interface{}) {
		fType := field.Type()

		var err error

		switch fType.Kind() {
		case reflect.Int:
			val, e := strconv.Atoi(value.(string))
			value = val
			err = e
		case reflect.Float32:
			val, e := strconv.ParseFloat(value.(string), 32)
			value = val
			err = e
		case reflect.Float64:
			val, e := strconv.ParseFloat(value.(string), 64)
			value = val
			err = e
		case reflect.Bool:
			val, e := strconv.ParseBool(value.(string))
			value = val
			err = e
		case reflect.Uint:
			val, e := strconv.ParseUint(value.(string), 10, 0)
			value = val
			err = e
		default:
			err = nil
		}
		if err == nil {
			vValue := reflect.ValueOf(value)
			field.Set(vValue)
		} else {
			fmt.Printf("%s\n", err.Error())
		}
	}

	names := strings.Split(fieldName, ".")
	currentObj := objValue

	count := len(names) - 1

	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("err from %s %v\n", fieldName, x)
		}
	}()
	for index, tempName := range names {
		if isArray, name, aIndex := getArray(tempName); isArray {
			field := currentObj.FieldByName(name)

			if field.Kind() != reflect.Slice {
				fmt.Printf("setField failed %s type error need slice\n", name)
				return
			}

			currentCount := field.Len() - 1

			if aIndex > currentCount {
				countDiff := aIndex - currentCount
				slice := reflect.MakeSlice(field.Type(), countDiff, countDiff)
				field.Set(reflect.AppendSlice(field, slice))
			}

			field = field.Index(aIndex)

			if index == count {
				setValue(field, value)
			} else {
				if field.IsNil() {
					field.Set(this.createInst(field.Type()))
				}
				if field.Kind() == reflect.Ptr {
					currentObj = field.Elem()
				} else {
					currentObj = field
				}
			}
		} else {
			if strings.ContainsAny(tempName, "[]") {
				fmt.Printf("setField failed <%s> format error\n", tempName)
				return
			}

			field := currentObj.FieldByName(tempName)
			if index == count {
				setValue(field, value)
			} else {
				if field.Kind() != reflect.Ptr {
					fmt.Printf("setField failed %s type error need Ptr\n", tempName)
					return
				}
				if field.IsNil() {
					inst := this.createInst(field.Type())
					field.Set(inst)
				}
				currentObj = field.Elem()
			}
		}

	}
}

func (this *ReflectUtil) createInst(objType reflect.Type) reflect.Value {
	newObj := reflect.New(objType.Elem())
	return newObj
}

func (this *ReflectUtil) typeSetFields(objType reflect.Type, values map[string]interface{}) interface{} {
	newObj := this.createInst(objType)

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
