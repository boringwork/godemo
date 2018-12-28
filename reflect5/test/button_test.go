package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/boringwork/godemo/reflect5"
)

func TestReflect(t *testing.T) {
	btn := &reflect5.Button{}

	typeBtn := reflect.TypeOf(btn)

	fmt.Println("package path: ", typeBtn.PkgPath())
	fmt.Println("name: ", typeBtn.Name())
}
