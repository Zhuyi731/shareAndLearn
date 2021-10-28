title: go语言学习分享
speaker: zoeizhu

<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

# go语言学习分享 {.text-landing.text-shadow}

<slide :class="size-40 ">
1. go简介 
2. Hello World
3. 基本类型和派生类型   
4. 数组、切片、函数、协程  
5. OOP
6. go语言的潜规则
7. 搭建一个简易http服务

<slide :class="size-30 aligncenter ">

!![](./imgs/go128.png .size-50.alignleft)

Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 {.text-intro.animated.fadeInUp.delay-500}  

Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。{.animated.fadeInUp}


<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

# go语言优点 {.text-landing.text-shadow}

<slide :class="size-30 aligncenter">

---
```shell {.animated.fadeInUp}
1、 学习曲线容易 （语法简洁、直来直去）

2、 效率  （接近C和C++的效率）

3、 出身名门、血统纯正

4、 强大的标准库

5、 部署方便：二进制文件，Copy部署

6、 语言层面支持并发 （协程）
```

<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

### go语言著名项目

<slide :class="size-30 aligncenter">

--- 

```shell{.animated.fadeInUp.delay-500}
1、 docker

2、 k8s

3、 gin
```


<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

## Hello World


<slide :class="size-40 aligncenter" >

---
hello.go文件
```go
package main  // 包名  同一个文件夹下包名相同

import "fmt" // 导入

func main() { // main包下的main函数是入口
    fmt.Println("Hello, World!")
}

```
<slide :class="size-60 topslide-top content-center">

## 基本类型

---
|            |                                                                                                                                                            |
| ---------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 布尔型     | 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。                                                                                |
| 数字类型   | 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。                                                      |
| 字符串类型 | 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。                 |
| 派生类型   | 包括：(a) 指针类型（Pointer）  (b) 数组类型  (c) 结构化类型(struct)  (d) Channel 类型  (e) 函数类型  (f) 切片类型  (g) 接口类型（interface）  (h) Map 类型 |


<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

## 基本类型变量声明

<slide :class="size-50">

--- 

```go
var a int // 整形还包括 uint8 uint16 uint32 uint64 int8~64 
a = 8 
b := 9 

var a float = 123.45 // 浮点类型同样包括float32 float64  complex32 complex64 


var c string 
c = "123"  
d := 'e' // rune类型  

var d bool = true 
e := false 

```
<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

## 派生类型变量声明  

<slide :class="size-50 lefttop" style="padding-top:0">

--- 

```go
var a = [3]int{1,2,3} // 数据大小为3的数组 
a := []int{1,2,3} // 容量为3长度为3的切片  
b = &a  // b为a的指针  

type Books struct {
   title string
   author string
   subject string
   book_id int
}  // 结构体声明

book1 := Books{"闺蜜之主", "爱潜水的乌贼", "克苏鲁", 0}
book2 := Book{ 
   title: "闺蜜之主",
   author: "爱潜水的乌贼",
   subject: "克苏鲁",
   book_id: 1
}  // 结构体赋值 

mapper := map[string]string  // map类型
mapper["123"] = "321"
fmt.Println(mapper["123"])
// 还有chan和interface没有说明   因为这两个使用比较复杂 会在后面说明
```

<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

# 数组和切片的区别 {.text-landing}

见demo

<slide :class="size-50 lefttop " style="padding-top:0">
# **go语言中的函数** {.text-landing.text-shadow.lightSpeedIn.animated.slow}
go语言函数是一等公民，使用方式和js非常相似{.lightSpeedIn.animated.slow}
支持闭包 匿名函数等

```go
// 函数声明  参数声明   返回值声明 
func sum(a int,b int) (c int){
    c = a + b // c已经定义
    return c 
}

// 匿名函数 
func (){
    fmt.Println("this is an anonymous function")
}()

// 闭包  
func closure() func(){
    ct := 100 
    return func(){
        ct = ct + 100 
        fmt.Printf("current ct is %s/n", ct)
    }
}
```


<slide class="bg-black aligncenter" image="https://cn.bing.com/az/hprichbg/rb/RedAntarctica_EN-AU12197122155_1920x1080.jpg">

## go OOP


<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">
## go中的对象


<slide :class="size-50 aligncenter">
---
go并不支持class关键字  

```go
type Animal struct {
    Kind string
    Age int
}

type Cat struct {
    Animal
}

type Dog struct {
    Animal
}
```

<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

# go 协程（goroutine）

<slide>

指在后台中运行的轻量级执行线程，go 协程是 Go 中实现并发的关键组成部分。  

Go 中提供了一个关键字 go 来让我们创建一个 Go 协程，当我们在函数或方法的调用之前添加一个关键字 go，这样我们就开启了一个 Go 协程，该函数或者方法就会在这个 Go 协程中运行。

```go
package main 

import "fmt"

func main(){
    fmt.Println("main start")

    go func(){
        fmt.Println("go routine")
    }()

    fmt.Println("main end")
}
```

<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">
# go语言中的潜规则  

go 语言中有内置的规则，如果不对这些进行了解可能会出现一些意向不到的错误。{.text-intro}

<slide class="aligncenter">

```shell
1. 同一文件夹下包名必须一致  

2. 导入模块时init函数会自动运行

3. 导入模块时前面加_表示只运行init函数，可以重命名

4. main包的main函数是入口  

5. struct和包中大写为public变量（函数）小写为private变量（函数）

6. go有统一的format格式

7. 只要实现了接口中声明的函数，就视为继承了该接口（隐式继承）
```


<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

# 使用go搭建一个简易http服务  

<slide class="bg-black-blue aligncenter" image="./imgs/bg.jpg .dark">

# The End