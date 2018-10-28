package main

import (
	"fmt"
	"reflect"
)

type People interface {
	Study()
}

type Student struct {
}

func (stu *Student) Study() {
	fmt.Println("stu.Study")
}

type Mine struct {
	Student
}

func (me *Mine) Study() {
	fmt.Println("Mine.Study")
}

func main() {
	appnil := (*Mine)(nil)

	appType := reflect.TypeOf(appnil)
	fmt.Println("app type: ", appType)
	app := appType.Elem()
	fmt.Println("appType elem: ", app)
	fmt.Println(app.Name())
	fmt.Println(app.PkgPath())
	mine := reflect.New(app).Elem().Interface()
	// p := app.(People)
	fmt.Println(mine)
	if p, ok := mine.(People); ok == true {
		fmt.Println(p)
	} else {
		fmt.Println("不行")
	}

	var peo People = &Mine{}
	fmt.Println(peo)
	if p, ok := peo.(People); ok == true {
		fmt.Println(p)
	} else {
		fmt.Println("不行")
	}
	// fmt.Println(m)
}
