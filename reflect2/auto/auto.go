package auto

import (
	"github.com/boringwork/godemo/reflect2/impl"
	"github.com/boringwork/godemo/reflect2/stu"
)

func Manifest(a int) stu.IStudent {
	if a == 0 {
		return &impl.MainActivity{}
	} else {
		return &impl.HomeActivity{}
	}
}
