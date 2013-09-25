package server

import (
	"fmt"
	"testing"
)

func TestPushGet(t *testing.T) {
	sc := new(SiteCollection)
	site := new(Site)
	sc.Push(site)
	if len(sc.array) != 1 {
		t.Fail()
	}

	ts, err := sc.Get(0)
	if err != nil {
		t.Logf("error:%s", err.Error())
		t.Fail()
	}
	if ts == nil {
		t.Fail()
	}
	if ts != site {
		t.Fail()
	}
	fmt.Printf("TestPushGet Pass\n")
}

func TestInsert(t *testing.T) {
	sc := new(SiteCollection)
	site := new(Site)
	sc.Push(new(Site))
	sc.Insert(0, site)

	ts, err := sc.Get(0)
	if err != nil {
		t.Logf("error:%s", err.Error())
		t.Fail()
	}
	if ts == nil {
		t.Fail()
	}
	if ts != site {
		t.Fail()
	}

	fmt.Printf("TestInsert Pass\n")
}
