package main

/*
#include "stdio.h"

#ifdef GODEBUG
int a = 1;
#else
int a = 2;
#endif

void* give(void* data){
	printf("%d\n", (int)data);
	return data;
}
*/
import "C"

import (
	"fmt"
	"runtime/debug"
	"unsafe"
)

type San struct {
	Name string
}

func main() {
	yusan := &San{Name: "雨伞"}

	up := unsafe.Pointer(yusan)

	C.give(up)
	fmt.Printf("a = %d\n", C.a)
	debug.PrintStack()
}
