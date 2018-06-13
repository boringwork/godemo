package impl

import (
	"fmt"

	"github.com/boringwork/godemo/interface/sss"
)

type BoyStudent struct {
	sss.Student
}

func (me *BoyStudent) study() {
	fmt.Println("BoyStudent study")
}

type GirlStudent struct {
	sss.Student
}

func (me *GirlStudent) study() {
	fmt.Println("GirlStudent study")
}
