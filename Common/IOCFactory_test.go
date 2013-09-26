package Common

import (
	"log"
	"reflect"
	"testing"
)

func RegistTest(t *testing.T) {
	factory := GetIOCFactory()
	factory.Regist(&testInterface, &reflect.TypeOf(testType), InstanceType_Normal) interface{}))

}

type testType struct {
}

type testInterface interface {
	Test()
}

func (this *testType) Test() {
	log.Printf("Test method run")
}
