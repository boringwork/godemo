package app

import (
	"fmt"
)

type IActivity interface {
	onStart()
}

type Activity struct {
}

func (me *Activity) onStart() {
	fmt.Println("Activity onStart")
}

type ActivityShadow struct {
	activity IActivity
}

func (me *ActivityShadow) onStart() {
	me.activity.onStart()
}

type activityManager struct {
	activities map[int]*ActivityShadow
	mainifest  func(int) IActivity
}

var am = &activityManager{
	activities: map[int]*ActivityShadow{},
}

func ActivityManager() *activityManager {
	return am
}

func (me *activityManager) SetManifest(manifest func(int) IActivity) {
	me.mainifest = manifest
}

func (me *activityManager) GetActivity(index int) *ActivityShadow {
	shadow, ok := me.activities[index]

	// TODO:: mux
	if !ok {
		act := me.mainifest(index)
		shadow = &ActivityShadow{activity: act}

		me.activities[index] = shadow
	}
	return shadow
}

func (me *activityManager) Run() {
	for i := 0; i != 2; i++ {
		me.GetActivity(i).onStart()
	}
}
