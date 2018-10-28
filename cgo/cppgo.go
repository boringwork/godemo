package main

/*
#cgo CFLAGS:  -I.
#cgo LDFLAGS:  -lstdc++
void _showAge();
*/
import "C"

func main() {
	C._showAge()
}
