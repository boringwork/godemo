package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Todo struct {
	Width   int
	Height  int
	Text    string
	Gravity string

	OnClick func(string)
}

func (me Todo) OnTap(v string) {
	fmt.Println("OnTap: ", v)
}

func TestTemplate(t *testing.T) {
	todo := &Todo{}
	// todo.OnTap("123")
	// a := reflect.ValueOf(todo).Elem()
	// // c := a.FieldByName("OnTap")
	// // b := a.Field(4)
	// // fmt.Println(a, c.Kind(), b.Kind())

	// fmt.Println("NumField")
	// for i := 0; i != a.NumField(); i++ {
	// 	f := a.Field(i)
	// 	fmt.Printf("%s, %s\n", a.Type().Field(i).Name, f.Kind())
	// }

	// fmt.Println("NumMethod")
	// for j := 0; j != a.NumMethod(); j++ {
	// 	f := a.Method(j)
	// 	fmt.Printf("%s, %s\n", a.Type().Method(j).Name, f.Kind())
	// 	// v := make([]reflect.Value, 0)
	// 	arg0 := "hello"
	// 	v := []reflect.Value{reflect.ValueOf(arg0)}
	// 	f.Call(v)
	// }

	// fmt.Println("========")

	callMethod(todo, "OnTap", "hello")
}

func callMethod(value interface{}, method string, args ...interface{}) {
	//TODO:: check value is a pointer
	a := reflect.ValueOf(value).Elem()
	f := a.FieldByName(method)
	if f.Kind() == reflect.Invalid {
		f = a.MethodByName(method)
	}

	if f.Kind() == reflect.Invalid || f.Kind() != reflect.Func {
		panic("method not found")
	}

	fmt.Println(args[0])
	b := make([]reflect.Value, 0)
	for _, arg := range args {
		// TODO:: check value or reference
		// 如何正好对应上参数
		b = append(b, reflect.ValueOf(arg))
	}
	f.Call(b)
}
