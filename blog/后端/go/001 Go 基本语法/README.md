# Go 快速入门基础掌握基本语法



###  1. 经典hello word

先了解一下一个最简单的 go 文件包含哪些内容

```go
// 包声明，每个 .go 文件必须声明
package main
// 引入依赖库
import "fmt"
// main 入口函数
func main() {
  // 使用 fmt 库的 Println 方法，打印 hello world
  fmt.Println("Hello world！")
}
```

运行：使用命令行运行或使用 ide 运行

```shell
$ go run hello.go
```

编译：Go是一门编译型语言，Go语言的工具链将源代码及其依赖转换成计算机的机器指令，生成可执行文件。

```shell
$ go build hello.go			// 生成可执行文件 hello
$ ./hello 							// 运行编译后的程序
Hello world！						// 输出
```



### 2. 内建基础变量类型



| 类型名称 | 类型                                                   | 初始值（zero value） |
| -------- | ------------------------------------------------------ | -------------------- |
| 整型     | (u)int、(u)int8、(u)int16、(u)int32、(u)int64、uintptr | 0                    |
| 浮点数   | float32、float64                                       | 0                    |
| 复数     | complex64、complex128                                  | 0                    |
| 布尔值   | bool                                                   | false                |
| 字符串   | string                                                 | ""                   |

* 整型：string 
* 浮点数：

### 2. 变量





### 1.2 变量

#### 1.2.1 变量声明

变量声明使用 `var` 关键字或使用简短声明方式 `:=` 。

```go
// 声明语句如下
var 变量名字 类型 = 表达式
```

> 其中 “类型” 跟 “= 表达式” 是可以省略的其中之一的，如果类型省略，编译器会根据表达式，自动推断，如果表达式省略，编译器会给变量赋初始值（zero value）



```go
package main

import "fmt"
// 包作用域下不能使用 := 声明变量，只能使用 var 声明

// 定义一个 foo 变量，类型为 string，默认初始值为 ""
var foo srting
// 定义一个 bar 变量，值为 "bar"，自动推断为 string 类型
var bar = "bar"

// 只定义，没有赋值。go 语言会默认给变量赋初始值（zero value）
func variableZeroValue() {
  // 定义一个类型为 int 的变量 a，默认赋值是 0
  var a int
  // 定义一个类型为 string 的变量 s，默认赋值是 ""
  var s string
  // 打印变量，%d 打印数值变量，使用 %q 可以打印出 ""
  fmt.Printf("%d %q\n", a, s) 
}
// 定义并赋初始值
func variableInitialValue() {
  // 定义一个类型为 int 的变量 a，并且赋值为 1
  var a int = 1
  // 同时定义 b，c 两个类型都为 int 的变量，并且分别赋值为 2，3
  var b, c int = 2, 3
  // 定义一个变量 d，并且赋值为 4。编译器会自动根据 4 推断类型为 int
  var d = 4
  // 声明变量的类型为 string，但是赋值为 5（int），类型不匹配会报错
  // var e string = 5 // 报错
  // 定义一个类型为 string 的变量 s，并赋值为 'abcde'
  var s string = 'abcde'
  
  fmt.Printfln(a, b, c, d ,s)
}

func main() {
  variableZeroValue()
  variableInitialValue()
}
```



总结：

使用var关键字

* `var a, b, c bool `

* `var s1, s2 string = "hello", "world"`

* 可放在函数内，可放在包内

* 使用 `var ()` 集中定义变量

让编译器自动决定类型

* `var a, b, i, s1, s2 = true, false, 3, "hello", "world"`

 使用 := 定义变量

* `a, b, i, s1, s2 := true, false, 3, "hello", "world"`



### 内建变量类型

* bool，string
* (u)int、(u)int8、(u)int16、(u)int32、(u)int64、uintptr

带u 的是无符号整数，不带u 的事有符号整数

有符号整数分两类，一类是规定长度的 int8、int16、int32、int64，一类是不规定长度的 int，根据系统是32位它就是32位，系统是64 位，它就是 64位。

最后一个 unitptr 中 ptr 是指针的意思，长度跟操作系统

* byte，rune

byte，代表字节，长度 8位

rune 是go 语言的 char 类型（字符型变量），rune 的长度是 32 位的，unicode 两个字节（16位），utf-8 很多字符都是 3 字节，所以 go 采用 int32

这两种类型可以跟整数混用，相当于整数的 别名。

* float32、float64、complex64、complex128

float32、float64 指定长度的浮点数

complex64、complex128 指定长度的 复数，有实部跟虚部，complex 64 的实部和虚部分别是 32位，complex128 的实部跟虚部分别是 64 位



complex

i = 根号-1

复数：z=a+bi，a 是实部，b是虚部，例如：3+4i

$|3+4i| = \sqrt{3^2 + 4^2} = 5$ 

$i^2 = -1$, $i^3 = -i$, $i^4 = i$

欧拉公式：$e^{ix} = cosx + isinx$



#### 强制类型转换

类型转换是强制的

```go
var a, b int = 3, 4
// c = math.Sqrt(a*a + b*b) // 报错，类型不匹配，必须显式转换
c = int(math.Sqrt(float64(a*a + b*b))
```

float类型是不准确的，有可能最后计算出来的结果不是5，是4.999999..然后转成整形变成4。

怎么解决？



### 常量

常量使用 const 关键字声明

常量的数值可以做各种类型使用（不明确指定类型的情况下）

常量命名不必全部大写，go语言大小写有其他含义（其他部分语言常量可能使用全部大写）

```go
const (
  filename = "abc.txt"
  a, b     = 3, 4
)
// a,b 不指定类型，就当作一个普通文本，下面计算的时候，不用转成 float64，但如果指定 a,b 是 int 就需要转
var c int
c = int(math.Sqrt(a*a + b*b))
fmt.Println(filename, c)
```



#### 枚举

iota 关键字

```go
const(
  one = iota
  _
  three
  four
)
const (
  b = 1 << (10 * iota)
  kb
  mb
  gb
  tb
  pb
)
```





#### 变量定义要点回顾

变量类型写在变量之后

编译器可以推测变量类型

没有char 只有 rune(32位)

原生支持复数类型





### 条件语句

#### if

```go
if v > 100 {
  return 100
} else if v < 0 {
  return 0
} else {
  return
}
```

* if 条件里不需要括号

```go
const filename = "abc.txt"
if contents, err := ioutil.ReadFile(filename); err != nil {
  fmt.Println(err)
} else {
  fmt.Printf("%s\n", contents)
}
```

* if 条件里可以跟一个赋值表达式
* if 条件里赋值的变量的作用域就在这个 if 语句里

#### switch

```go
switch op {
  cate "+":
  	result = a + b
  cate "+":
    result = a + b
  cate "+":
    result = a + b
  cate "+":
    result = a + b
  default:
    panic()
}
switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
}
```

* switch 会自动 switch， 除非使用 fallthrough
* switch 后可以没有表达式，在 case 里面加判断条件



### 循环

#### for

```
sum := 0
for i := 1; i <= 100; i++ {
  sum += i
}
```

* for 条件里不需要括号
* for 的条件里可以省略初始条件、结束条件、递增条件

* 没有 while， for 可以当 while 使用



### 函数

go 语言中，函数是一等公民

```go
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}
```

* 函数使用 fanc 关键字定义
* 函数返回类型写在最后面
* 函数可以作为函数的参数（函数式编程）
* 函数可以返回多个值，常用场景是第二个返回值用于返回错误信息

```go
func div(a, b int) (q, r int) {
  q := a/b
  r := a%b
  // 这里直接 return 会返回 q, r，代码比较多的时候，不建议这样做，代码会不清晰
	return
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

q, r := div(13, 3)

```

* 函数可以给返回值定义名字
* ~~部分编辑器可以自动根据函数调用生成返回值名称~~
* 没有默认参数，没有可选参数
* 函数没有重载
* 有可变参数列表

```go
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
```



### 指针

go 语言的指针比较简单

```go
var a int = 2
var pa *int = &a
*pa = 3
fmt.Println(a)
```

* 指针不能运算（c 可以）



#### 参数传递

参数传递一般有 值传递或引用传递 两种传递方式

go 语言只有值传递一种方式

普通值传递：

```go
var a int
func f(a int)

f(a)
```

指针传递：利用指针传递实现类似引用传递的效果

```go
var a int 
func f(pa *int)

f(&a)
```


```go
var cache Cache
func f(cache Cache)

f(cache)
```

cache 是一个指向data 的指针，传递的时候，传递cache 的值。

cache 一般缓存比较多数据，如果数据都保存在 cache 变量里，直接传递，会耗费比较多资源。

go 语言的自定义的类型在定义的时候，就要考虑到这种情况，用这种类型的时候，是用它的指针呢，还是用它的值。这里 cache 就是用它的值。



一个交换变量值的例子：

```go
func swap(a, b int) {
  b, a = a, b
}
func main() {
  a, b := 3, 4
  fmt.Println(swap(a,b))
}
```

使用这种方式无效，因为参数传递的是值的拷贝，不会影响原来的参数

```go
func swap(a, b *int) {
  *b, *a = *a, *b
}
func main() {
  a, b := 3, 4
  fmt.Println(swap(&a, &b))
}
```

这种方式可以交换成功，因为使用指针传递，参数传递的是值的内存地址，函数内再通过内存地址取值改值，就能成功改掉传递进去的参数的值。



```
func swap(a, b int) (int, int) {
	return b, a
}
```





### 数组

```go
var array1 [5]int
array2 := [3]int{1, 2, 3}
array3 := [...]int{1, 2, 3}
var grid = [4][5]int
```

* 数量写在类型前面

遍历数组：

```go
// 一般我们可能想到这样的方式
for i := 0; i < len(array3); i++ {
  fmt.Println(array3[i])
}
// 但 go 语言有更简洁的方式
for i := range array3 {
  fmt.Println(array3[i]) // 打印的结果是一样的
}
for i, v := range array3 {
  fmt.Println(array3[i], v)
}
for _, v := range array3 {
  fmt.Println(v)
}
```

使用 range 可以获取数组的下标跟 值，值可以不取，可以通过` _` 省略变量

不仅在 range，在任何地方都可以使用 `_` 省略变量



#### 为什么要用 range

* 意义明确、美观

#### 数组是值类型

* 函数传递数组，会传数组的值的拷贝

* 方括号内的值不同或方括号后的类型不同，都表示是不同的类型。

  例如: `[5]int` 和 `[3]int` 是不同的类型

* Go 语言一般不会直接使用数组，而是使用切片（slice）



### 切片（slice）

```go
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s := arr[2:6]
// s = [2, 3, 4, 5]
```

切片是数组的一个视图

```go
s1 := arr[2:6]
s2 := arr[2:]
s3 := arr[:6]
s4 := arr[:]
```

切片是[x:y] 这样的结构，x，y可以省略

* slice 本身是没有数据的，是对底层 array 的一个 view
* slice 的值改变，会修改对应的数组的值

reslice

```go
s1 = s1[:3] // [2, 3, 4]
s1 = s1[1:] // [3, 4]
s1 = s1[:] // [3, 4]
```



#### slice 的扩展

```go
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6] // [2, 3, 4, 5]
s2 : = s1[3:5] // ? s1[3], s1[4]
```

s1 是 [2, 3, 4, 5]，只有 4 个值，s2 取 s1 的 3到5，超出了s1 的范围，但是它还是可以取出来，返回的是 [5, 6]

![](./images/slice-extendsion.png)

![](./images/slice-extendsion2.png)

 slice 有 3 属性。ptr 指针，len 长度，cap 容量

slice 可以向后扩展，不能向前扩展

s[i] 不能超越 len(s)，向后扩展不能超越底层数组cap(s)

#### 向 slice 添加元素

```go
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6] // [2, 3, 4, 5]
s2 : = s1[3:5]
s3 := append(s2, 10)
s4 := append(s3, 11)
s5 := append(s4, 12)
```

* 添加元素时，如果未超越原数组cap，则会改变原数组对应位置的值

* 添加元素时，如果超过了原数组cap，系统会重新分配更大的底层数组，原来的数组如果有人用，会继续存在，如果没有人用，则垃圾回收机制会回收该数组

* 由于值传递的关系，必须接收 append 的返回值 `s = append(s, val)`

#### 复制 slice

```go
// 把s1 拷贝给 s2
copy(s2, s1)
```



#### 删除 slice 中某个元素

```go
// 删除索引为 3 的元素
s2 = append(s2[:3], s2[4:]...)
// Popping from front 删除第一个元素
front := s2[0]
s2 = s2[1:]
// Popping from back 删除最后一个元素
tail = s2[len(s2) - 1]
s2 = s[:len(s2) - 1]
```



### Map

```go
m := map[string]string {
  "name": "silin"
  "language": "golang"
}
m2 := make(map[string]int) // m2 == empty map
var m3 = map[string]int // m3 == nil
```

* 普通 map：`map[K]V`，复合 map：`map[K1]map[K2]V` 



#### map 遍历

#### map 创建

```go
for k rang m {
  fmt.Println(k);
}
for k, v rang m {
  fmt.Println(k, v);
}
for _, v rang m {
  fmt.Println(v);
}
```

* key 是无序的，每次输出的顺序可能不同，

#### map 取值

```go
name := m["name"]
fmt.Println(name)
// 取一个不存在的key 的值，返回map 值的类型的 zero value
// 这里取的 age 不存在，m 的值的类型是 string，所以返回的 zero value 是 ""，空字符串
// age := m["age"]
// 可以通过第二值返回值判断 key 是否存在，存在返回 true， 不存在返回 false
if age, ok := m["age"]; ok {
  fmt.Println(age)
} else {
  fmt.Println("key does not exist")
}
```



#### map 删除值

```go
name, ok = m["name"]
fmt.Println(name, ok) // silin true
// 删除 m 里面的 "name"
delete(m, "name")
name, ok = m["name"]
fmt.Println(name, ok) // "" false
```



* 使用 len 函数获取元素个数
* 排序需要手动排序，可以把key 放到 slice 里，再通过 slice 排序



#### map 的key

* map 使用 hash 表，必须可以比较相等
* 除了 slice、map、function 的内建类型，其他类型都可以作为key（slice、map、function 不能比较相等）
* Struct 类型不包含以上类型，也可以作为 key



#### 例子：寻找最长不含有重复字符的字符串

https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

```go

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
```



### 字符串
#### rune

Go 中没有 char，rune 相当于char

```go
func main() {
	s := "hello，世界!" // utf-8编码，每个中文 3 个字节
	fmt.Printf("%s\n", []byte(s))
	fmt.Printf("%X\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	fmt.Println(len(bytes))
	// 使用 utf8.DecodeRune 解码，循环打印每个字符
	for len(bytes) > 0 {
		// 返回首个字符以及字符占用字节大小
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c ", ch)
		// 去掉首个字符
		bytes = bytes[size:]
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
```



* 使用 range 遍历 pos, rune 对

#### 其他字符串操作

* 库 strings，在操作字符串的时候，首先先看 strings 库有没有内置相关方法

* 常用方法：

  Fields，Split，Join

  Contains，Index

  Tolower，Toupper

  Trim，TrimRight，TrimLeft



### 面向对象

go 语言只支持封装，不支持继承与多态（继承与多态用接口来做）

go 语言没有 class，只有 struct



#### struct（结构）的定义

使用 type 关键字定义一个类型为 struct 的 TreeNode

```go
type treeNode struct {
	Value int
  Left, Right *treeNode,
}
```

下面是创建赋值的几种方法示例

```go
var root treeNode
root = treeNode{Value: 3}
root.Left = &treeNode{}
root.Right = &treeNode{5, nil, nil}
root.Right.Left = new(treeNode)

// 在 slice 的定义里，可以省略每个子元素的 trueNode 类型
nodes := []treeNode {
  {Value: 3}, // 相当于 {Value: 3, Left: nil, Right: nil}
  {}, // 相当于 {Value: 0, Left: nil, Right: nil}
  {6, nil, nil} // 相当于 {Value: 6, Left: nil, Right: nil}
}
```

* 不论是指针地址还是结构本身，一律使用 . 来访问成员
* Go语言没有构造函数，struct 没有构造函数，使用 `treeNode{}` 或 `&treeNode{}` 或 `new(treeNode)` 来构造。
* 或创建工厂函数，然后通过工厂函数进行构造

```go
// 工厂函数
func createNode(value int) *treeNode {
  return &treeNode{value: value}
}
```

```go
root.Left.Right = createTreeNode(2)
```

* 工厂函数返回了局部变量（闭包）

#### 这个局部变量的结构创建在堆上还是栈上？

在 Go 语言中，不必知道是要分配在堆上还是栈上，编译器会自动分配。如果函数中一个变量，没有取地址，没有返回到函数外使用，那编译器可能就认为这个变量不需要给外面用，就会在栈上分配。如果是取地址并且返回给函数外面用，编译器可能就会在堆上分配，如果外面拿到指针，使用完后不用这个指针了，编译器就会进行垃圾回收处理。

#### 为结构定义方法

```go
func (t treeNode) print() {
  fmt.Print(t.value)
}
// 如果需要修改接受器本身，则需要使用指针接收者
func (t *treeNode) setValue(value int) {
  t.value = value
}
```

在方法前面加一个括号，括号里是一个接收者，通过这种方式把方法定义到 struct 上

```go
// 使用 struct 的方法
root.print()
root.setValue(5)
```

* 编译器会自动识别接收者是指针接收者还是值接收者，不需要在调用的时候指明。
* 接收器类似于其他语言的 this 或者 self，但go语言不指定要用 this 或 self，可以自定义变量名。但一般情况下是使用 struct 名称的第一个字母的小小写。例如上面 struct 是 treeNode，所以用 t。
* nil 指针也可能调用方法

```go
func (t *treeNode) setValue(value int) {
  if t == nil {
    fmt.Println("Setting value to nil node. Ignored.")
    return
  }
  t.value = value
}

var nilRoot *treeNode // nilRoot = nil
nilRoot.setValue(100) // Print "Setting value to nil node. Ignored."
```



例子：遍历上面定义的树节点

```go
func (t *treeNode) traverse() {
  if t == nil {
    return
  }
  t.left.traverse()
  t.print()
  t.right.traverse()
}
```



#### 值接收者 vs 指针接收者

* 要改变内容必须使用指针接收者
* 结构过大也考虑使用指针接收者
* 一致性：如有指针接收者，最好全部统一定义为指针接收者
* 值接收者 是go 语言特有（php？）
* 值／指针接收者均可接收值／指针



#### 封装

* 名字一般使用 CamelCase（驼峰）
* 首字母大写：public（公有方法）
* 首字母小写：private（私有方法）



#### 包

* 每个目录一个包
* main 包包含可执行入口
* 为结构定义的方法必须放在同一个包
* 可以是不同的文件

#### 如果扩充系统类型或者别人的类型

* 定义别名
* 使用组合

##### 使用组合

````go
// 在外面包一层
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}
func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	fmt.Print("My own post-order traversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
````

##### 定义别名

```go
package queue
type Queue []int
func (q *Queue) Push(v int) {
  *q = append(*q, v)
}
func (q *Queue) Pop() int {
  head := (*q)[0]
  *q = (*q)[1:]
  return head
}
func (q *Queue) IsEmpty() bool {
  return len(*q) == 0
}
```



```go
package main

import "path/to/queue"
func main() {
  q := queue.Queue{1}
  q.Push(2)
  q.Push(3)
  fmt.Println(q.Pop())
  fmt.Println(q.Pop())
  fmt.Println(q.IsEmpty())
  fmt.Println(q.Pop())
  fmt.Println(q.IsEmpty())
}

```

> 这里定义 Push、Pop 会改变 q 本身



### GOPATH 环境变量

默认在 ~/go (unix, linux), %USERPROFILE%\go (windows)

官方推荐：所有项目和第三方库都放在同一个 GOPATH 下

也可以将不同的项目放在不同的 GOPATH



#### go get 获取第三方库

`go get xxx`



goimport



#### gopath 下目录结构

* src 项目文件
* pkg
* bin 可执行文件



#### 包引入

```go
import (
	"fmt"
	"path/to/第三方包"
)
```



##### 包

每个文件夹下只能有一个 main包跟一个 main 函数



## 接口

```go
type Traversal interface {
  Traverse()
}
func main () {
  traveral := getTraversal()
  traveral.Traverse()
}
```



### duck typing

##### 大黄鸭是鸭子吗？

* 传统类型系统：脊索动物门，脊椎动物亚门，鸟纲雁形目...
* duck typing：是鸭子，因为长得像鸭子
* ”像鸭子走路，像鸭子叫，长得像鸭子，那么就是鸭子“
* 描述事物的外部行为而非内部结构
* 严格说 go 语言属于结构化类型系统，类似 duck typing。
* duck typing 说需要动态绑定，但是 go 语言是编译时绑定



##### go 语言的 duck typing

同时需要 Readable，Appendable 怎么办？（java：apache polygene）

同时具有 python，c++ 的duck typing 的灵活性

又具有 java 的类型检查



##### 接口定义

download（使用者）-》retriever（实现者）

接口由使用者定义

*main.go*

```go
package main

import (
	"fmt"
  "path/to/retriever"
)

// 定义 Retriever 接口，接口有 GET 方法
type Retriever interface {
	Get(url string) string
}

// 使用者
func download(r Retriever) string {
	return r.Get("https://www.silinchen.com")
}

func main() {
	var r Retriever
	r = retriever.Retriever{}
	fmt.Println(download(r))
}

```

retriever.go

```go
package retriever

import (
	"net/http"
	"net/http/httputil"
	"time"
)
// 定义一个 Retriever 结构，这里定义为什么名字都行
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}
// 实现一个 Get 方法
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
```



##### 接口的实现

* 接口的实现是隐式的
* 



```go
func main() {
	var r Retriever
	r = retriever.Retriever{
    UserAgent: "Mozilla/5.0",
    TimeOut: time.Minute,
  }
	fmt.Printf("%T %v\n", r, r) // retriever.Retriever {Mozilla/5.0 1m0s}
  // r 有类型和值
}
```



使用指针

path/to/retriever.go

```go
func (r *Retriever) Get(url string) string {
  ...
}
```
main.go
```go
func main() {
	var r Retriever
	r = retriever.Retriever{
    UserAgent: "Mozilla/5.0",
    TimeOut: time.Minute,
  }
	fmt.Printf("%T %v\n", r, r) // *retriever.Retriever &{Mozilla/5.0 1m0s}
  // r 有类型和值
}
```



##### 接口变量里面有什么

接口变量有：实现者的类型，实现者的值或指向实现者的指针

接口变量自带指针

接口变量同样采用值传递，几乎不需要使用接口的指针



##### 查看接口变量

* 表示任何类型：interface{}
* Type Assertion x.(T)
* Type Switch



##### 接口的组合

```go
type Poster interface {
  Post(url srting, form map[string]string) string
}
func post(poster Poster) {
  poster.Post("https://silinchen.com",
              map[string]string {
                "name": "silin",
                "site": "silinchen.com"
              }
  )
}
type RetriverPoster interface {
  Retriver
  Poster
}
func session() {
  
}
```



## 函数式编程

函数与闭包

```go
func adder() func(value int) int {
  sum := 0
  return func(value int) int {
     sum += value
     return sum
  }
}
func main() {
  adder := adder()
  for i := 0; i < 10; i++ {
    fmt.Println(adder(i))
  }
}
```



* 函数是一等公民：参数，变量，返回值都可以是函数
* 高阶函数
* 函数 -》闭包



#### “正统” 函数式编程

* 不可变性：不能有状态，只有常量和函数
* 函数只能有一个参数
* 本课程不做上述严格规定

上面例子改为“正统 ”函数式写法

```go
type iAdder func(int) (int, iAdder)
func adder2(base int) iAdder {
  return func(v int) (int, iAdder) {
    return base + v, adder2(base + v)
  }
}
```

* 这种写法比较晦涩，不直观



#### 斐波那契数列

```go
package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func main() {
	f := fibonacci()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 5
	fmt.Println(f()) // 8
  ...
}

```

#### 为函数实现接口

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}
// 定义一个函数的类型 intGen，并作为上面函数的返回类型
type intGen func() int
// 给函数定义方法
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// 如果 p 太小，会出问题，需要改进
	return strings.NewReader(s).Read(p)
}
func printContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printContents(f)
}

```



#### 遍历二叉树

高阶函数

```go
func (n *Node) TraverseFunc (f func (n*Node)) {
	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}
```



### Go 语言闭包的应用

* 更为自然，不需要修饰如何访问自由变量
* 没有 Lambda 表达式，但是有匿名函数



## 错误处理和资源管理

#### 资源管理

##### defer 调用

* 确保调用在函数结束时发生

```go
defer fmt.Println(1)
defer fmt.Println(2)
fmt.Println(3)
// panic("error")
return
fmt.Println(4)
// 输出结果：
// 3
// 2
// 1
```



defer 在函数结束的时候执行

defer 里面相当于一个栈，栈是先进后出的，所以先声明的后执行

```go
func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
  // 关闭已打开的文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 将缓存中的内容写到文件中
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
```



* 参数在 defer 语句时计算

  ```go
  for i:=0;i<=30;i++ {
    defer fmt.Print(i, " ")
  }
  // 输出:30 29 28 ... 3 2 1
  // 而不是输出 30 个 30
  ```

* defer 列表为后进先出



#### 何时使用 defer 调用

* Open／Close
* Lock／Unlock
* PrintHeader／PrintFooter



#### 错误处理

```go
file, err := os.Open("abc.txt")
if err := nil {
  if pathError, ok := err.(*os.PathError); ok {
    fmt.Println(pathError.Err)
  } else {
    fmt.Println("unknown error", err)
  }
}
```

> 针对已知类型跟未知类型分别处理
>
> 这样的处理比较复杂，繁琐

如何实现统一的错误处理？



#### 服务器统一出错处理

errhanling/filelistingserver/filelisting/hanlder.go

```go
func HandleFileList(writer http.ResponseWriter, request *http.Request)) error {
  path := request.URL.Path[len("/list/"):]
  file, err := os.Open(path)
  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)
    return err
  }
  defer file.Close()
  all, err := ioutil.ReadAll(file)
  if err != nil {
    return err
  }
  writer.Write(all)
  return nil
}
```

errhanling/filelistingserver/web.go

```go
type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
  err := handler(writer, request)
  if err != nil {
    log.Warn("Error handling request: %s", err.Error())
    code := http.StatusOK
    switch {
      // 404 不存在
      case os.IsNotExist(err):
      	code = http.StatusNotFound
      // 403 没权限
      case os.IsPermission(err):
        code = http.StatusForbidden
      default:
        code = http.StatusInternalServerError
    }
    http.Error(writer, http.StatusText(code), code)
  }
}
func main() {
  http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
  err := http.ListenAndServe(":8888", nil)
  if err != nil {
    panic(err)
  }
}
```



#### panic

* 停止当前函数执行
* 一直向上返回，执行每一层的 defer
* 如果没有遇见recover，程序退出



#### recover

* 仅在 defer 调用中使用
* 获取 panic 的值
* 如果无法处理，可重新 p











