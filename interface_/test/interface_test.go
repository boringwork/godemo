package test

import (
	"testing"

	"github.com/boringwork/godemo/interface_"
)

func TestInterface(t *testing.T) {
	mainAct := &interface_.MainActivity{}

	OnCreate, ok := interface{}(mainAct).(interface_.ActivityOnCreate)
	if ok {
		OnCreate.OnCreate()
	} else {
		t.Fail()
	}

	OnDestroy, ok := interface{}(mainAct).(interface_.ActivityOnDestroy)
	if ok {
		OnDestroy.OnDestroy()
	} else {
		t.Fail()
	}
}
