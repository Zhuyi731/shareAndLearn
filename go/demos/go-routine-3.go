package main

import (
	"fmt"
	"sync"
)

func ctSum(upper int, w *sync.WaitGroup, c chan int) int {
	sum := 0
	for i := 0; i < upper; i++ {
		sum = sum + i
	}
	// fmt.Println(sum)
	w.Done()
	c <- sum
	return sum
}

var w sync.WaitGroup

// 主协程与子协程通信
func main() {
	c := make(chan int)

	w.Add(3)
	go ctSum(100, &w, c)
	go ctSum(10, &w, c)
	go ctSum(1000, &w, c)

	result1, result2, result3 := <-c, <-c, <-c
	fmt.Println(result1, result2, result3)
	w.Wait()
}
