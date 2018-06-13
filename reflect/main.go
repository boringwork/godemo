package main

import (
	"github.com/boringwork/godemo/reflect/app"
	"github.com/boringwork/godemo/reflect/auto"
)

func main() {
	app.ActivityManager().SetManifest(auto.Manifest)
	app.ActivityManager().Run()
}
