package stu

import (
	"fmt"
)

type IStudent interface {
	study()
}

type Student struct {
}

func (me *Student) study() {
	fmt.Println("Activity onStart")
}

type StudentShadow struct {
	student IStudent
}

func (me *StudentShadow) study() {
	me.student.study()
}

type activityManager struct {
	activities map[int]*StudentShadow
	mainifest  func(int) IStudent
}

var am = &activityManager{
	activities: map[int]*StudentShadow{},
}

func ActivityManager() *activityManager {
	return am
}

func (me *activityManager) SetManifest(manifest func(int) IStudent) {
	me.mainifest = manifest
}

func (me *activityManager) GetActivity(index int) *StudentShadow {
	shadow, ok := me.activities[index]

	// TODO:: mux
	if !ok {
		act := me.mainifest(index)
		shadow = &StudentShadow{student: act}

		me.activities[index] = shadow
	}
	return shadow
}

func (me *activityManager) Run() {
	for i := 0; i != 2; i++ {
		me.GetActivity(i).study()
	}
}
