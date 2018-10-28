package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/boringwork/godemo/class/stu"
)

type MyType struct {
	i    int
	name string
}

func (mt *MyType) SetI(i int) {
	mt.i = i
}

func (mt *MyType) SetName(name string) {
	mt.name = name
}

func (mt *MyType) String() string {
	return fmt.Sprintf("%p", mt) + "--name:" + mt.name + " i:" + strconv.Itoa(mt.i)
}

func test() {
	myType := &MyType{22, "wowzai"}
	//fmt.Println(myType)     //就是检查一下myType对象内容
	//println("---------------")
	mtV := reflect.ValueOf(&myType).Elem()
	fmt.Println("Before:", mtV.MethodByName("String").Call(nil))
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	mtV.MethodByName("SetI").Call(params)
	params[0] = reflect.ValueOf("reflection test")
	mtV.MethodByName("SetName").Call(params)
	fmt.Println("After:", mtV.MethodByName("String").Call(nil))
	f := mtV.MethodByName("SetI")
	fmt.Printf("%T\n", f)
}

func main() {
	boy := &stu.BoyStudent{}
	mtV := reflect.ValueOf(&boy).Elem()
	fmt.Println("settability of v:", mtV.CanSet())
	// m := mtV.MethodByName("Study")
	// fmt.Printf("%T\n", m)
	// m.Call(nil)

	// myType := &MyType{22, "wowzai"}
	// mtV = reflect.ValueOf(&myType).Elem()
	// typeOfT := mtV.Type()
	fmt.Println(mtV.NumMethod())
	// for i := 0; i != mtV.NumMethod(); i++ {
	// 	mtV.Field(i)
	// 	// fmt.Printf("%d: %T\n", i, f)
	// 	// fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	// }
	// test()
}
