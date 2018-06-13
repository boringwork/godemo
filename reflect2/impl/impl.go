package impl

import (
	"fmt"

	"github.com/boringwork/godemo/reflect2/stu"
)

type MainActivity struct {
	stu.Student
}

func (me *MainActivity) study() {
	fmt.Println("*MainActivity onStart")
}

type HomeActivity struct {
	stu.Student
}

func (me *HomeActivity) study() {
	fmt.Println("*HomeActivity onStart")
}
