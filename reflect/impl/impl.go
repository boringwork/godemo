package impl

import (
	"fmt"

	"github.com/boringwork/godemo/reflect/app"
)

type MainActivity struct {
	app.Activity
}

func (me *MainActivity) onStart() {
	fmt.Println("*MainActivity onStart")
}

type HomeActivity struct {
	app.Activity
}

func (me HomeActivity) onStart() {
	fmt.Println("HomeActivity onStart")
}
