package main

import "fmt"

/**
数组和切片的差别
**/
func main() {
	// part1
	// go中所有基础类型的拷贝都是值拷贝
	var a = [5]int{}
	var b = [5]int{}
	for i := 1; i <= 5; i++ {
		a[i-1] = i
	}

	fmt.Println(a)
	fmt.Println(b)

	b = a // 赋值
	b[0] = 6

	fmt.Println(a)
	fmt.Println(b)

	// part 2
	// 切片存的是数组的地址引用
	var c = make([]int, 5)
	var d = make([]int, 5)
	for i := 1; i <= 5; i++ {
		c[i-1] = i
	}
	fmt.Println(c)
	fmt.Println(d)
	d = c // 赋值
	d[0] = 6

	fmt.Println(c)
	fmt.Println(d)
}
