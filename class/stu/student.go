package stu

import "fmt"

type IStudent interface {
	study()
}

type Student struct {
}

func (me *Student) study() {
	fmt.Println("Student study")
}

type BoyStudent struct {
	A int
	*Student
}

func (me *BoyStudent) study() {
	fmt.Println("BoyStudent study")
}

func (me *BoyStudent) Study() {
	fmt.Println("BoyStudent study   SSS ")
}

func (me *BoyStudent) Jo() {
	fmt.Println("BoyStudent study from JO")
}

type GirlStudent struct {
	*Student
}

func (me *GirlStudent) study() {
	fmt.Println("GirlStudent study")
}
