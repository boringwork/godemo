package main

import (
	"github.com/boringwork/godemo/reflect2/auto"
	"github.com/boringwork/godemo/reflect2/stu"
)

func main() {
	stu.ActivityManager().SetManifest(auto.Manifest)
	stu.ActivityManager().Run()
}
