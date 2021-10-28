package main

import "fmt"

/**
数组和切片的差别
**/
func main() {
	b := make([]int, 1024*1024)

	fmt.Println(cap(b))
	for i := 1; i <= 1024*1024; i++ {
		b[i-1] = i
	}
	// fmt.Println(b)

	// var a = [5]int{}
	// for i := 1; i <= 5; i++ {
	// 	a[i-1] = i
	// }
	// fmt.Println(a)
	cb := cap(b)
	fmt.Println(cap(b))
	b = append(b, 6)
	fmt.Println(cap(b))
	fmt.Println(cap(b) - cb)

	// fmt.Println(b)
	// a = append(a, 6)
	// fmt.Println(a)
}
