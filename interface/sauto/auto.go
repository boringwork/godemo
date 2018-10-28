package sauto

import (
	"fmt"

	"github.com/boringwork/godemo/interface/aaa"
	"github.com/boringwork/godemo/interface/simpl"
)

func init() {
	fmt.Println("init package sauto")
}

func Finder(name string) aaa.IStudent {
	switch name {
	case "boy":
		return &simpl.BoyStudent{}
	case "girl":
		return &simpl.GirlStudent{}
	}
	return nil
}
