package main

/*
#cgo CXXFLAGS: -std=c++11

#include "callstu.h"
*/
import "C"

func main() {
	C.display()
}
