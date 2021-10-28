package main

import (
	"fmt"
)

// 另外一种方式   面向对象的函数
type Animal struct {
	Kind string
}

func (a *Animal) eat(food string) {
	fmt.Println("This is a " + a.Kind + " it's eating " + food)
}

type Cat struct {
	Animal
}

func (c *Cat) eat(food string) {
	// c.Kind="xxx"
	fmt.Println("Cat eat " + food)
}

func newCat(kind string) &Cat{
	return &Cat{
		Animal: Animal{
			Kind: kind,
		},
	}
}

func main() {
	// https://github.com/golang/go/issues/9859
	// go 语言层面的缺陷
	c := Cat{
		Animal: Animal{
			Kind: "cat",
		},
	}
	// c := Cat {
	// 	Kind: "xx",
	// }
	// c.Kind = "dog"

	c.eat("fish")
}
