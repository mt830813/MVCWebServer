package Util

import (
	"fmt"
	"testing"
)

func TestRunMethod(t *testing.T) {
	inst := new(ReflectUtil)

	tObj := new(a)

	if results, err := inst.RunDMethod(tObj.Test, "test"); err != nil {
		t.Logf("error:%s\n", err.Error())
		t.Fail()
	} else {
		result := results[0]
		if result != "testb" {
			t.Logf("result not right\n")
			t.Fail()
		}
		if oResults, e := inst.RunDMethod(tObj.Test2, tObj, "test"); err != nil {
			t.Logf("error:%s\n", e.Error())
			t.Fail()
		} else {
			oResult := oResults[0]
			if result != oResult {
				t.Logf("result not right \n")
				t.Fail()
			} else {
				fmt.Printf("%s\n", oResult)
			}
		}
	}

	if oResults, err := inst.RunDObjMethod(tObj, "Test", "test"); err != nil {
		t.Logf("error:%s\n", err.Error())
		t.Fail()
	} else {
		if results, e := inst.RunDMethod(tObj.Test, "test"); e != nil {
			t.Logf("error:%s\n", e.Error())
			t.Fail()
		} else {
			if results[0] != oResults[0] {
				t.Logf("result not right\n")
				t.Fail()
			}
		}
	}

}

func TestGetParam(t *testing.T) {
	reflect := new(ReflectUtil)

	param := map[string]interface{}{
		"Name": "Tom",
		"Sex":  "Male",
	}

	tObj := &a{Name: param["Name"].(string)}

	obj := new(a)

	if result, err := reflect.RunMapMethod(obj.Test3, param); err != nil {
		t.Log(err.Error())
		t.Fail()
		if result[0] != obj.Test3(tObj) {
			t.Log("result error")
			t.Fail()
		}
	}

}

type a struct {
	Name string
	Sex  string
}

func (this *a) Test(param string) string {
	return param + "b"
}

func (this *a) Test2(obj *a, param string) string {
	return obj.Test(param)
}

func (this *a) Test3(obj *a) string {
	return obj.Name
}
