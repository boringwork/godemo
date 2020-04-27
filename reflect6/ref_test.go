package reflect6

import (
	"fmt"
	"reflect"
	"testing"
)

type a interface {
	Sum()
}

type b struct {
}

func TestReflect6(t *testing.T) {
	at := reflect.TypeOf((*a)(nil)).Elem()
	fmt.Printf("%T, %s\n", at, at.PkgPath())
	bt := reflect.TypeOf((*b)(nil)).Elem()
	fmt.Printf("%T, %s\n", bt, bt.PkgPath())
}
