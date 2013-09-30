package Common

import (
	"fmt"

	"reflect"
	"testing"
)

func TestFactoryGet(t *testing.T) {
	factory := GetIOCFactory()
	factory2 := GetIOCFactory()

	if factory != factory2 {
		t.Logf("GetIOCFactory is not singleton")
		t.Fail()
	}
}

func TestRegist(t *testing.T) {

	var ti *testInterface
	bType := reflect.TypeOf(ti).Elem()
	factory := GetIOCFactory()
	factory.Regist(bType, reflect.TypeOf(new(testType)), InstanceType_Normal, nil)
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

	factory.Regist(bType, reflect.TypeOf(new(testType2)), InstanceType_Normal, nil)

	length = len(factory.getPArray(bType))

	if length != 1 {
		t.Logf("interface check function failed")
		t.Fail()
	}
}

func TestNormalInst(t *testing.T) {
	factory := GetIOCFactory()

	ti := reflect.TypeOf((*testInterface)(nil)).Elem()

	factory.Regist(ti, reflect.TypeOf(new(testType)), InstanceType_Normal, CreateTest)
	factory.RegistByName("other", ti, reflect.TypeOf(new(testType3)), InstanceType_Singleton, nil)

	rResult := CreateTest(2).Test()

	tObj, _ := factory.Get(ti, []interface{}{2})
	oObj, _ := factory.GetByName("other", ti, nil)
	oObjTwo, _ := factory.GetByName("other", ti, nil)
	tResult := tObj.(testInterface).Test()
	oResult := oObj.(testInterface).Test()

	fmt.Printf("")

	tTest := tObj.(*testType)

	if tTest.key != 2 {
		t.Logf("result err,key:%d", tTest.key)
		t.Fail()
	}

	if tResult != rResult {
		t.Logf("instance error\ntResult:%d rResult:%d", tResult, rResult)
		t.Fail()
	}
	if oResult == tResult {
		t.Logf("instance error\ntResult:%d oResult:%d", tResult, oResult)
		t.Fail()
	}

	if oObj != oObjTwo {
		t.Logf("instance error")
		t.Fail()
	}
}

func TestDecorateInst(t *testing.T) {
	factory := GetIOCFactory()

	ti := reflect.TypeOf((*testInterface)(nil)).Elem()

	factory.Regist(ti,
		reflect.TypeOf(new(testType)), InstanceType_Normal, nil)
	factory.RegistDecorate(ti, reflect.TypeOf(new(testTypeDecorater)))

	var rObj = new(testType)
	var rResult = rObj.Test()

	var tObj, _ = factory.Get(ti, nil)
	var tResult = tObj.(testInterface).Test()

	if tResult != rResult*4 {
		t.Logf("decorate result error")
		t.Fail()
	}

}

type testType struct {
	key int
}

type testType2 struct {
}

type testType3 struct {
}

type testInterface interface {
	Test() int
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

func (this *testType) Test() int {
	return this.key
}

func (this *testType3) Test() int {
	return 256
}

func CreateTest(key int) *testType {
	returnValue := new(testType)
	returnValue.key = key
	return returnValue
}
