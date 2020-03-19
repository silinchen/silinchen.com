package main

import "fmt"

// 顶级作用域声明变量只可以使用 var 关键词声明
var foo string = "foo"

// 不声明变量类型，go 语言会根据所赋的值的类型推断变量的类型
var bar = "bar"

// 顶级作用域不能使用 := 赋值
// bar := "bar" // 报错

// 括号语法
var (
	a = "a"
	b = "b"
	c = 3
)

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
	var a int = 1
	// 多个同类型的变量可以使用
	var b, c int = 2, 3
	// 声明类型为 string 但是赋值为数值会报错
	// var d string = 4 // 报错
	var s string = "abcde"
	// 有赋值初始化的，可以用 := 操作符赋值，这种方式不必声明类型，go 语言会自行推断类型
	foo := "foo"
	// 声明类型会报语法错误
	// bar string := "bar"  // 报错
	fmt.Println(a, b, c, s, foo)
}

// 变量类型推断
func variableTypeDeduction() {
	var a = 1
	var foo = "foo"
	fmt.Println(a, foo)
}

// 简短定义语法
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func main() {
	// 使用 fmt 库的 Println 方法，打印 hello world
	fmt.Println("Hello world！")
	variableZeroValue()
	variableInitialValue()
	variableShorter()
	fmt.Println(foo, bar, a, b, c)
}
