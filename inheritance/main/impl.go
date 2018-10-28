package main

import "github.com/boringwork/godemo/inheritance/aaa"

type U struct {
	*aaa.T
}

func (me *U) m() {
	println("FAIL")
}
