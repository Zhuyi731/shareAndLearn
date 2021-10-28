package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")
	a := 1
	go func() {
		a = 2
		fmt.Println("go routine")
	}()
	go func() {
		fmt.Println(a)
		fmt.Println("go routine")
	}()

	fmt.Println("main end")
	time.Sleep(10 * time.Millisecond)
}
