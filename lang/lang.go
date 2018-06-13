package main

import (
	"fmt"
)

func main() {
	over := false
	v := 0
	for {
		v = getv()
		switch v {
		case 5:
			fmt.Println("555555")
			over = true
			break
		}
		if over {
			break
		}
	}
	fmt.Println(v)
}

var index int = 0

func getv() int {
	return 5
	// index++
	// return index
}
