package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	test5()
}

func test5() {
	var x float64 = 3.4
	fmt.Println("x = ", x)
	p := reflect.ValueOf(&x) // 注意:得到X的地址 fmt.Println("type of p:", p.Type()) fmt.Println("settability of p:" , p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println("x = ", x)
}

func test4() {
	type Q struct {
		A int
		B string
	}

	type T struct {
		Q
		C float32
	}
	s := reflect.TypeOf(T{})
	fmt.Println(s.Field(0))
	v := reflect.New(s)
	// v.Interface
	fmt.Println(&v)
	a, ok := (&v).Interface().(*T)
	// a.B = "hello"
	fmt.Println(a, ok)
	fmt.Println(unsafe.Pointer(a), unsafe.Pointer(&a))
}

func test3() {
	type T struct {
		A int
		B string
	}
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func test2() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())
	v.SetFloat(4.4)
	fmt.Println("value:", v.Float())
	fmt.Println("value:", x)
}

func test1() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // 注意:得到X的地址 fmt.Println("type of p:", p.Type()) fmt.Println("settability of p:" , p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
