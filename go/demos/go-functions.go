package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("这是一个匿名函数")
	}()

	increaseCt := closure()
	ct := increaseCt()
	fmt.Println(ct)
	ct = increaseCt()
	fmt.Println(ct)
}

func closure() func() int {
	ct := 0
	return func() int {
		ct++
		return ct
	}
}
