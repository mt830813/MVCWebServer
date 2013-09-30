package Util

import (
	"testing"
)

func TestRunMethod(t *testing.T) {
	inst := new(ReflectUtil)

	tObj := new(a)

	if results, err := inst.RunDMethod(tObj.Test, "test"); err != nil {
		t.Logf("error:", err.Error())
		t.Fail()
	} else {
		result := results[0]
		if result != "testb" {
			t.Logf("result not right")
			t.Fail()
		}
	}

	if oResults, err := inst.RunDObjMethod(tObj, "Test", "test"); err != nil {
		t.Logf("error:", err.Error())
		t.Fail()
	} else {
		if results, e := inst.RunDMethod(tObj.Test, "test"); e != nil {
			t.Logf("error:", e.Error())
			t.Fail()
		} else {
			if results[0] != oResults[0] {
				t.Logf("result not right")
				t.Fail()
			}
		}
	}

}

type a struct {
}

func (this *a) Test(param string) string {
	return param + "b"
}
