// rules1 同一文件夹下的包名必须一致
package main

// rules2
import (
	// _ "github.com/xxx/xxx"
	// lru "github.com/go-lru"
	"fmt"
)

// rules3
func init() {
	fmt.Println("init function")
}

//rules4
type People struct { // 大写表示public成员  小写为private成员
	Name string `json: "name" form:"name"` // tag
	age  int
}

type action interface {
	say(string)
}

// rules5
func (p *People) say(something string) {
	fmt.Println(p.Name + " said:" + something)
}

// rules6 main包中的main函数为唯一入口
func main() {
	p := new(People)
	p.Name = "Biden"
	p.age = 16
	p.say("I'm newed")

	a := true
	//rules7 go有同一的格式化工具  必须严格按照其格式缩进  否则可能会出现意想不到的问题
	if a {
		fmt.Println("it's ok!")
	}

	// if a
	// {
	// 	fmt.Println("it cannot work!")
	// }
}
