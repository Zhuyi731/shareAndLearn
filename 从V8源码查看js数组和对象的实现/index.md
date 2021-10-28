# js中对象的实现  

看到这个标题，可能会比较疑惑。js中对象不是一个基本数据类型吗，哪有什么实现不实现的。   
我们可以横向对比一下其他的语言，可以发现许多的静态类型语言都没有对象这种数据类型（python 字典，rust 散列）。而我们一般用的js引擎v8又是用C++写的。  
所以中间肯定有某层实现来帮助我们方便的使用对象这种数据结构。   

## js中对象的特性   

先看一个例子   
猜一下哪些不能作为key值   

```js

var obj = {} 

const keys = [1,'abc',{},[],()=>{},new Map(),new Symbol('123'),new Set()]

keys.forEach(key=>{
  try {
    obj[key] = key
  } catch(err) {
    console.log(key)
  }
})
 
console.log(obj)  //?
```


```js  
var obj1 = {a:'aa'}
var obj2 = {b:'bb'}

var obj = {}
obj[obj1] = '1'
obj[obj2] = '2' 
console.log(obj)  // ?
```

```js  
var obj = {} 
obj[100] = 100
obj[1] = 1
obj[2] = 2 

obj['c'] = 'c' 
obj['a'] = 'a'
obj['b'] = 'b'

for(var key in obj){
    console.log(key) // 输出顺序
}

```