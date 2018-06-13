package sss

import (
	"fmt"
)

type IStudent interface {
	study()
}

type Student struct {
}

func (me *Student) study() {
	fmt.Println("Student study")
}

type StudentProxy struct {
	student IStudent
}

func (me *StudentProxy) study() {
	me.student.study()
}

type studentManager struct {
	students map[string]*StudentProxy
	finder   func(string) IStudent
}

var sm = &studentManager{
	students: map[string]*StudentProxy{},
}

func StudentManager() *studentManager {
	return sm
}

func (me *studentManager) SetFinder(finder func(string) IStudent) {
	me.finder = finder
}

func (me *studentManager) GetStudent(name string) *StudentProxy {
	proxy, ok := me.students[name]
	if !ok {
		stu := me.finder(name)
		proxy = &StudentProxy{student: stu}
		me.students[name] = proxy
	}
	return proxy
}

func (me *studentManager) MethodA() {
	names := []string{"boy", "girl"}
	for _, name := range names {
		me.GetStudent(name).study()
	}
}

func (me *studentManager) MethodB() {
	names := []string{"boy", "girl"}
	for _, name := range names {
		me.GetStudent(name).study()
	}
}
