package main

import (
	"fmt"
	"sync"
)

func ctSum(upper int, w *sync.WaitGroup) int {
	sum := 0
	for i := 0; i < upper; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
	w.Done()
	return sum
}

var w sync.WaitGroup

// 任务分配

func main() {

	w.Add(3)
	go ctSum(100, &w)
	go ctSum(10, &w)
	go ctSum(1000, &w)

	w.Wait()
}
