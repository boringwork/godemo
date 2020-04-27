package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	width := js.Global().Get("innerWidth")
	fmt.Printf("width type is %T\n", width)
	// println("renderFrame ", width)
	js.Global().Set("myint", 1)
	js.Global().Set("myfunc", js.FuncOf(myfunc))
	js.Global().Set("add2", js.FuncOf(add2))
	js.Global().Set("mystring", "my string")
	js.Global().Set("add", js.FuncOf(add))

	console_log := js.Global().Get("console").Get("log")
	console_log.Invoke("Hello wasm! console_log in go")

	js.Global().Call("eval", `
        console.log("hello, wasm!");
	`)

	js.Global().Call("eval", `
		console.log(add(1,2));
	`)

	js.Global().Set("a_bool", js.ValueOf(true))
	js.Global().Set("a_int", js.ValueOf(123))
	js.Global().Set("a_string", js.ValueOf("abc"))

	js.Global().Call("eval", `
        console.log(typeof a_bool, a_bool);
        console.log(typeof a_int, a_int);
        console.log(typeof a_string, a_string);
	`)

	js.Global().Call("eval", `
	a_bool = true;
	a_int = 123;
	a_string = "abc";
`)

	println(js.Global().Get("a_bool").Bool())
	println(js.Global().Get("a_int").Int())
	println(js.Global().Get("a_string").String())
}

func myfunc(this js.Value, args []js.Value) interface{} {
	println("run go func 'myfunc'")
	return nil
}

// function definition
func add(this js.Value, i []js.Value) interface{} {
	println("call add ", i[0].Int()+i[1].Int())
	return js.ValueOf(i[0].Int() + i[1].Int())
}

func add2(this js.Value, args []js.Value) interface{} {
	s1 := js.Global().Get("document").Call("getElementById", "v1").Get("value").String()
	s2 := js.Global().Get("document").Call("getElementById", "v2").Get("value").String()
	v1, _ := strconv.Atoi(s1)
	v2, _ := strconv.Atoi(s2)
	println(v1 + v2)
	return v1 + v2
}

// exposing to JS
