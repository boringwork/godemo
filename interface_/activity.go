package interface_

import (
	"fmt"
)

type ActivityOnCreate interface {
	OnCreate()
}

type ActivityOnDestroy interface {
	OnDestroy()
}

type MainActivity struct {
	Render func()
}

func (me *MainActivity) OnCreate() {
	fmt.Println("MainActivity OnCreate")
}

func (me *MainActivity) OnDestroy() {
	fmt.Println("MainActivity OnDestroy")

}
