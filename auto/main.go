package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

// mArr = []string{"mleft", "center", "mright", "wheelDown",
// "wheelUp", "wheelLeft", "wheelRight"}

var open = false

func main() {
	go func() {
		for {
			ok := robotgo.AddEvent("0")
			if ok {
				fmt.Println("pressed 0")
				open = true
			}
		}
	}()

	go func() {
		for {
			ok := robotgo.AddEvent("+")
			if ok {
				fmt.Println("pressed +")
				open = false
			}
		}
	}()

	go func() {
		for {
			fmt.Println("KeyTap f ?")
			if open {
				fmt.Println("KeyTap f")
				robotgo.KeyTap("f")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	robotgo.AddEvent("center")
}
