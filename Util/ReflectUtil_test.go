package Util

import (
	"fmt"
	"reflect"
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
	rUtil := new(ReflectUtil)

	param := map[string]interface{}{
		"Name":   "Tom",
		"Sex":    "1",
		"A.Name": "Jim",
		"A.Sex":  "2",
	}

	tObj := &a{Name: param["Name"].(string)}

	obj := new(a)

	if result, err := rUtil.RunMapMethod(obj.Test3, param); err != nil {
		t.Log(err.Error())
		t.Fail()
	} else {
		if result[0] != obj.Test3(tObj) {
			t.Logf("result error %s\n", result[0])
			t.Fail()
		}
	}

}

func TestObjSetFields(t *testing.T) {
	rUtil := new(ReflectUtil)

	param := map[string]interface{}{
		"Name":      "Tom",
		"Sex":       "1",
		"A.Name":    "Jim",
		"A.Sex":     "2",
		"B[0].Name": "Green",
		"B[0].Sex":  "1",
		"C[0]":      "1",
		"C[1]":      "2",
	}
	obj := new(a)
	rUtil.ObjSetFields(reflect.ValueOf(obj).Elem(), param)
	if obj.Name != param["Name"] {
		t.Logf("result error %s\n", obj.Name)
		t.Fail()
	}
	if obj.A == nil {
		t.Logf("result error %v\n", obj.A)
		t.Fail()
	}
	if obj.A.Name != param["A.Name"] {
		t.Logf("result error %s\n", obj.A.Name)
		t.Fail()
	}
	if len(obj.B) != 1 {
		t.Logf("result error %v\n", obj.B)
		t.Fail()
	}
	if obj.B[0].Name != param["B[0].Name"] {
		t.Logf("result error %s\n", obj.B[0].Name)
		t.Fail()
	}
	if obj.C[0] != param["C[0]"] {
		t.Logf("result error %d\n", obj.C[0])
		t.Fail()
	}
	if obj.C[1] != param["C[1]"] {
		t.Logf("result error %d\n", obj.C[1])
		t.Fail()
	}
}

type a struct {
	Name string
	Sex  int
	A    *a
	B    []*a
	C    []string
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
