package main

/*
#include "stdio.h"

void* give(void* data){
	printf("%d\n", (int)data);
	return data;
}
*/
import "C"

import (
	"unsafe"
)

type San struct {
	Name string
}

func main() {
	yusan := &San{Name: "雨伞"}

	up := unsafe.Pointer(yusan)

	C.give(up)
}
