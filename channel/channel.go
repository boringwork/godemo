package main

import (
	"fmt"
	"time"
)

var ch chan int

func Count() {
	// fmt.Println("count 1")
	time.Sleep(1 * time.Second)
	fmt.Println("count 1")

	ch <- 1
	fmt.Println("count 2", "ch = ", ch)
	// time.Sleep(1 * time.Second)
	ch <- 2
	fmt.Println("count 3", "ch = ", ch)
}

func main() {
	test()

	if true {
		return
	}

	fmt.Println("start....")
	ch = make(chan int, 1)
	go Count()

	//fmt.Println("main 1")
	//time.Sleep(2 * time.Second)
	fmt.Println("main 1")
	<-ch
	fmt.Println("main 2", "ch = ", ch)
	time.Sleep(1 * time.Second)
	fmt.Println("main 3", "ch = ", ch)
	<-ch
	time.Sleep(2 * time.Second)
	fmt.Println("main 4", "ch = ", ch)
}

func test() {
	ch := make(chan int, 8)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			fmt.Print("number ")
			i := <-ch
			fmt.Printf("i = %d\n", i)
		}
	}()

	for i := 0; i != 10; i++ {
		fmt.Println("发送 ", i)
		ch <- i
	}

	fmt.Println("end")
}
