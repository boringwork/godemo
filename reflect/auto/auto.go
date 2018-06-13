package auto

import (
	"github.com/boringwork/godemo/reflect/app"
	"github.com/boringwork/godemo/reflect/impl"
)

func Manifest(a int) app.IActivity {
	if a == 0 {
		return &impl.MainActivity{}
	} else {
		return &impl.HomeActivity{}
	}
}
