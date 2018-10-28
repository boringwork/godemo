package simpl

import (
	"fmt"

	"github.com/boringwork/godemo/interface/aaa"
)

func init() {
	fmt.Println("init package simpl")
}

type BoyStudent struct {
	*aaa.Student
}

func (me *BoyStudent) study() {
	fmt.Println("BoyStudent study")
}

type GirlStudent struct {
	*aaa.Student
}

func (me *GirlStudent) study() {
	fmt.Println("GirlStudent study")
}
