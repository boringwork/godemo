package main

import (
	"encoding/json"
	"fmt"
)

type Activity interface {
	OnCreate()
	onDestroy()
}

type MainActivity struct {
	Title string
}

func (this *MainActivity) OnCreate() {
	fmt.Println("MainActivity" + "OnCreate")
}

func (this *MainActivity) onDestroy() {
	fmt.Println("MainActivity" + "onDestroy")
}

func onCreate(act Activity) {
	act.OnCreate()
	act.onDestroy()
}

func change(act *MainActivity) {
	act.Title = "你好"
	fmt.Println(act.Title)
}

type HomeActivity struct {
	MainActivity
}

type Integer int

func (me *Integer) Add1(v Integer) Integer {
	return *me + v
}

func main() {
	mainact := &MainActivity{}
	mainact.OnCreate()

	onCreate(mainact)

	act := MainActivity{}
	onCreate(&act)

	json.Marshal(act)
	json.Unmarshal([]byte(""), act)

	change(&act)
	fmt.Println("after: ", act.Title)

	var a Integer = 1
	b := &a
	fmt.Println(a.Add1(2))
	fmt.Println(b.Add1(2))

	home := &HomeActivity{}
	home.OnCreate()
}
