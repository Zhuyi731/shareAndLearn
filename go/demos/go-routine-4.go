package main

import (
	"fmt"
	"time"
)

// 通信阻塞

func main() {
	c := make(chan int)
	defer close(c)
	go func() {
		c <- 3 + 4
		fmt.Println("sended")
	}()
	// i := <-c
	time.Sleep(time.Second * 5)
	fmt.Println("ok")
}
