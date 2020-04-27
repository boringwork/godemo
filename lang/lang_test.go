package main

import (
	"fmt"
	"testing"
)

func TestLang(t *testing.T) {
	// over := false
	// v := 0
	// for {
	// 	v = getv()
	// 	switch v {
	// 	case 5:
	// 		fmt.Println("555555")
	// 		over = true
	// 		break
	// 	}
	// 	if over {
	// 		break
	// 	}
	// }
	// fmt.Println(v)

	fmt.Printf("%.2f", 1.0)
}

var index int = 0

func getv() int {
	return 5
	// index++
	// return index
}
