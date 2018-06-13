package main

import (
	"github.com/boringwork/godemo/interface/auto"
	"github.com/boringwork/godemo/interface/sss"
)

func main() {
	sss.StudentManager().SetFinder(auto.Finder)
	sss.StudentManager().MethodA()
}
