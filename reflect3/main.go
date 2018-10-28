package main

import (
	"fmt"
	"reflect"

	"github.com/boringwork/godemo/class/stu"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	typeRegistry[t.Name()] = t
	fmt.Println(t.PkgPath() + "/" + t.Name()) // echo github.com/boringwork/godemo/class/stu/Student
}

func newStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true
}

func init() {
	registerType((*test)(nil))
	registerType((*stu.Student)(nil))

}

type test struct {
	Name string
	Sex  int
}

type test2 struct {
	test
	age int
}

func main() {
	structName := "test2"

	// s1 := reflect.TypeOf(stu.Student)

	s, ok := newStruct(structName)
	if !ok {
		fmt.Println("没有")
		return
	}

	fmt.Println(s, reflect.TypeOf(s))

	t, ok := s.(test2)
	if !ok {
		return
	}
	t.Name = "i am test"
	fmt.Println(t, reflect.TypeOf(t))
}
