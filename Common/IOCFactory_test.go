package Common

import (
	"fmt"

	"reflect"
	"testing"
)

func TestRegist(t *testing.T) {

	var ti *testInterface
	bType := reflect.TypeOf(ti).Elem()
	factory := GetIOCFactory()
	factory.Regist(bType, reflect.TypeOf(new(testType)), InstanceType_Normal)
	length := len(factory.getPArray(bType))
	if length != 1 {
		t.Logf("Regist Test Failed")
		t.Fail()
	}
	context, err := factory.getRegistContext("default", bType)
	if err != nil {
		t.Logf("Regist Test Failed:%s", err.Error())
		t.Fail()
	}

	obj := context.(*registContext)

	if obj.bType != reflect.TypeOf(new(testType)) {
		t.Logf("bType error")
		t.Fail()
	}

	if obj.instType != InstanceType_Normal {
		t.Logf("instance type error")
		t.Fail()
	}

	factory.Regist(bType, reflect.TypeOf(new(testType2)), InstanceType_Normal)

	length = len(factory.getPArray(bType))

	if length != 1 {
		t.Logf("interface check function failed")
		t.Fail()
	}
}

func TestNormalInst(t *testing.T) {
	factory := GetIOCFactory()

	factory.Regist(reflect.TypeOf((*testInterface)(nil)).Elem(),
		reflect.TypeOf(new(testType)), InstanceType_Normal)
	factory.RegistByName("other", reflect.TypeOf((*testInterface)(nil)).Elem(),
		reflect.TypeOf(new(testType3)), InstanceType_Normal)

	rResult := new(testType).Test()

	ti := reflect.TypeOf((*testInterface)(nil)).Elem()

	tObj, _ := factory.Get(ti)
	oObj, _ := factory.GetByName("other", ti)
	tResult := tObj.(testInterface).Test()
	oResult := oObj.(testInterface).Test()
	tContext, _ := factory.getRegistContext("default", ti)
	rContext, _ := factory.getRegistContext("default", ti)
	oContext, _ := factory.getRegistContext("other", ti)

	fmt.Printf("%v,%v", tContext == oContext, rContext == tContext)
	if tResult != rResult {
		t.Logf("instance error\ntResult:%d rResult:%d", tResult, rResult)
		t.Fail()
	}
	if oResult == tResult {
		t.Logf("instance error\ntResult:%d oResult:%d", tResult, oResult)
		t.Fail()
	}

}

type testType struct {
}

type testType2 struct {
}

type testType3 struct {
}

type testInterface interface {
	Test() int
}

func (this *testType) Test() int {
	return 128
}

func (this *testType3) Test() int {
	return 256
}
