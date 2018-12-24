package test

import (
	"fmt"
	"testing"

	"github.com/boringwork/godemo/interface_"
)

func TestInterface(t *testing.T) {
	mainAct := &interface_.MainActivity{
		Render: func() {
			fmt.Println("render widget")
		},
	}

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

	mainAct.Render()
}
