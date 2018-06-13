package auto

import (
	"github.com/boringwork/godemo/interface/impl"
	"github.com/boringwork/godemo/interface/sss"
)

func Finder(name string) sss.IStudent {
	switch name {
	case "boy":
		return &impl.BoyStudent{}
	case "girl":
		return &impl.GirlStudent{}
	}
	return nil
}
