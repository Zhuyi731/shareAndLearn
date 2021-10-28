package inherit

import (
	"fmt"
)

type Abstract_Animal interface {
	eat(string)
	run()
}

type Animal struct {
	Kind string
}

func (a *Animal) eat(food string) {
	fmt.Println("This is a " + a.Kind + " it's eating " + food)
}

func (a *Cat) run(){

}

type Cat struct {
	Animal
}

func (c *Cat) eat(food string) {
	fmt.Println("Cat eat " + food)
}

func main() {
	// https://github.com/golang/go/issues/9859
	// go 语言层面的缺陷
	c := Cat{
		Animal: Animal{
			Kind: "cat",
		},
	}
	// c.Kind = "dog"

	c.eat("fish")
}
